package main

import (
	"axox/gaingauge/views"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	password  string
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
	router.Get("/favicon.ico", faviconHandler)
	router.Post("/register", env.registerHandler)
	router.Post("/login", env.loginHandler)

	router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(env.authentication)
		r.Get("/", indexHandler)
		// router.Get("/api", env.GetUsers)
	})
	http.ListenAndServe(":3000", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int64)
	templ.Handler(views.Index(strconv.FormatInt(userId, 10))).ServeHTTP(w, r)
}

func (env Env) authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("session_token")
		if err != nil {
			templ.Handler(views.Login()).ServeHTTP(w, r)
			return
		}
		sessionToken := c.Value

		var userId int64
		var expiresAt time.Time
		err = env.db.QueryRow(context.Background(), "select user_id, expires_at from loginsessions where id=$1 and now() < expires_at", sessionToken).Scan(&userId, &expiresAt)

		if err != nil {
			templ.Handler(views.Login()).ServeHTTP(w, r)
			return
		}

		if expiresAt.Before(time.Now()) {
			templ.Handler(views.Login()).ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (env Env) registerHandler(w http.ResponseWriter, r *http.Request) {

}

func (env Env) loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	var userId int64
	err := env.db.QueryRow(context.Background(), "select id from users where username = $1 and password = $2", username, password).Scan(&userId)
	if err != nil {
		w.Write([]byte(`Failed to log in`))
	}
	w.Header().Add("HX-Refresh", "true")
	expiresAt := time.Now().Add(2 * time.Minute)
	var sessionId string
	err = env.db.QueryRow(context.Background(), "insert into loginsessions (\"user_id\", \"expires_at\") values (1, $1) RETURNING id;", expiresAt).Scan(&sessionId)
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   sessionId,
		Expires: expiresAt,
	}
	http.SetCookie(w, cookie)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./resources/favicon.ico")
}
