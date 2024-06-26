package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Env struct {
	db *pgxpool.Pool
}

type User struct {
	id         int64
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
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

	router := gin.Default()
	router.GET("/users", env.getUsers)

	router.Run("localhost:8080")
}

func (e *Env) getUsers(c *gin.Context) {

	rows, err := e.db.Query(context.Background(), "select * from users")
	if err != nil {
		c.AbortWithStatus(500)
	}
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.id, &user.First_name, &user.Last_name); err != nil {
			c.AbortWithStatus(500)
		}
		users = append(users, user)
	}

	fmt.Println(len(users))
	c.JSON(http.StatusOK, users)
}

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }
//
// func postAlbums(c *gin.Context) {
// 	var newAlbum album
//
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
//
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
//
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")
//
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }
