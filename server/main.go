package main

import (
	"axox/gaingauge/views"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Env struct {
	db *pgxpool.Pool
}

func main() {
	godotenv.Load()

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	env := &Env{db: conn}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./resources/favicon.ico") })
	router.Get("/spinner.svg", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./resources/spinner.svg") })
	router.Get("/index.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./views/index.css") })
	router.Post("/register", env.registerHandler)
	router.Get("/register", templ.Handler(views.RegisterForm()).ServeHTTP)
	router.Get("/login", templ.Handler(views.LoginForm()).ServeHTTP)
	router.Post("/login", env.loginHandler)

	router.Group(func(r chi.Router) {
		r.Use(env.authentication)
		r.Get("/", indexHandler)
		r.Get("/log", env.logHandler)
		r.Get("/exercises", env.exercisesHandler)
		r.Post("/createWorkout", env.createWorkout)
		r.Post("/createExercise", env.createExercise)
		r.Get("/workoutEditor", env.workoutEditor)
		r.Get("/exerciseOverview", env.exerciseOverview)
	})
	http.ListenAndServe(":3000", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	isHtmxReq := r.Header.Get("HX-Request") == "true"
	w.Header().Add("Vary", "HX-Request")
	templ.Handler(views.Index(isHtmxReq)).ServeHTTP(w, r)
}

func (env Env) exercisesHandler(w http.ResponseWriter, r *http.Request) {
	isHtmxReq := r.Header.Get("HX-Request") == "true"
	w.Header().Add("Vary", "HX-Request")
	id, ok := r.Context().Value("userId").(int64)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	rows, err := env.db.Query(context.Background(), "select name from exercises where user_id = $1 order by name asc", id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var exercises []string
	for rows.Next() {
		var exercise string
		err := rows.Scan(&exercise)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		exercises = append(exercises, exercise)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	templ.Handler(views.Exercises(isHtmxReq, exercises)).ServeHTTP(w, r)
}

func (env Env) logHandler(w http.ResponseWriter, r *http.Request) {
	isHtmxReq := r.Header.Get("HX-Request") == "true"
	w.Header().Add("Vary", "HX-Request")
	id, ok := r.Context().Value("userId").(int64)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	rows, err := env.db.Query(context.Background(), "select date from workouts where user_id = $1 order by date desc", id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var dates []time.Time
	for rows.Next() {
		var date time.Time
		err := rows.Scan(&date)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		dates = append(dates, date)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	templ.Handler(views.Log(isHtmxReq, dates)).ServeHTTP(w, r)
}

func (env Env) authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHtmxReq := r.Header.Get("HX-Request") == "true"
		c, err := r.Cookie("session_token")
		if err != nil {
			if isHtmxReq {
				w.Header().Add("HX-Refresh", "true")
				return
			}
			templ.Handler(views.Login()).ServeHTTP(w, r)
			return
		}
		sessionToken := c.Value

		var userId int64
		var expiresAt time.Time
		err = env.db.QueryRow(context.Background(), "select user_id, expires_at from loginsessions where id=$1 and now() < expires_at", sessionToken).Scan(&userId, &expiresAt)

		if err != nil {
			if isHtmxReq {
				w.Header().Add("HX-Refresh", "true")
				return
			}
			templ.Handler(views.Login()).ServeHTTP(w, r)
			return
		}

		if expiresAt.Before(time.Now()) {
			if isHtmxReq {
				w.Header().Add("HX-Refresh", "true")
				return
			}
			templ.Handler(views.Login()).ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (env Env) registerHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")

	if username == "" {
		w.Write([]byte(`Please provide a nonempty username`))
		return
	}

	if password == "" {
		w.Write([]byte(`Please provide a nonempty password`))
		return
	}

	var count int64
	err := env.db.QueryRow(context.Background(), "select count(*) from users where username = $1", username).Scan(&count)
	if err != nil {
		w.Write([]byte(`Failed to register`))
		return
	}
	if count != 0 {
		w.Write([]byte(`Username already taken`))
		return
	}

	h := sha256.New()
	h.Write([]byte(password))
	hash := hex.EncodeToString(h.Sum(nil))
	var userId int64

	err = env.db.QueryRow(context.Background(), `insert into users ("username", "first_name", "last_name", "password") VALUES ($1, $2, $3, $4) returning id`, username, firstName, lastName, hash).Scan(&userId)
	if err != nil {
		w.Write([]byte(`Failed to register`))
		return
	}
	expiresAt := time.Now().Add(8 * time.Hour)
	var sessionId string
	err = env.db.QueryRow(context.Background(), "insert into loginsessions (\"user_id\", \"expires_at\") values ($1, $2) RETURNING id;", userId, expiresAt).Scan(&sessionId)
	if err != nil {
		w.Write([]byte(`Failed to login`))
		return
	}
	w.Header().Add("HX-Refresh", "true")
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   sessionId,
		Expires: expiresAt,
	}
	http.SetCookie(w, cookie)
}

func (env Env) createWorkout(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02T15:04", r.FormValue("date"))
	if err != nil {
		w.Header().Add("HX-Retarget", "#create-workout-error")
		w.Header().Add("HX-Reswap", "innerHTML")
		w.Write([]byte(`Invalid date format error`))
		return
	}

	id, ok := r.Context().Value("userId").(int64)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	_, err = env.db.Exec(context.Background(), `insert into workouts ("user_id", "date") values ($1, $2)`, id, date)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("HX-Trigger", "new-workout")
	w.Write([]byte(fmt.Sprintf(`<option value="%s" selected> %s </option>`, date.Format("2006-01-02T15:04"), date.Format("2006-01-02 15:04"))))
}

func (env Env) createExercise(w http.ResponseWriter, r *http.Request) {
	exercise := r.FormValue("exercise")
	if exercise == "" {
		w.Header().Add("HX-Retarget", "#create-exercise-error")
		w.Header().Add("HX-Reswap", "innerHTML")
		w.Write([]byte(`Please provide an exercise name`))
		return
	}

	id, ok := r.Context().Value("userId").(int64)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	_, err := env.db.Exec(context.Background(), `insert into exercises ("user_id", "name") values ($1, $2)`, id, exercise)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("HX-Trigger", "new-exercise")
	w.Write([]byte(fmt.Sprintf(`<option value="%s" selected> %s </option>`, exercise, exercise)))
}

func (env Env) exerciseOverview(w http.ResponseWriter, r *http.Request) {
	exercise := r.FormValue("exercise")
	w.Write([]byte(exercise))
}

func (env Env) workoutEditor(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02T15:04", r.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Write([]byte(date.String()))
}

func (env Env) loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	h := sha256.New()

	h.Write([]byte(password))

	hash := hex.EncodeToString(h.Sum(nil))

	var userId int64
	err := env.db.QueryRow(context.Background(), "select id from users where username = $1 and password = $2", username, hash).Scan(&userId)
	if err != nil {
		w.Write([]byte(`Failed to log in`))
		return
	}
	w.Header().Add("HX-Refresh", "true")
	expiresAt := time.Now().Add(8 * time.Hour)
	var sessionId string
	err = env.db.QueryRow(context.Background(), "insert into loginsessions (\"user_id\", \"expires_at\") values ($1, $2) RETURNING id;", userId, expiresAt).Scan(&sessionId)
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   sessionId,
		Expires: expiresAt,
	}
	http.SetCookie(w, cookie)
}
