package main

import (
"github.com/gin-gonic/gin"
// "net/http"
    "database/sql"
    "fmt"
  _  "log"
_    "net/http"
_ "github.com/mattn/go-sqlite3"
)

func main() {
    router := gin.Default()
    db := NewDB()
    router.GET("/", func(c *gin.Context) {
        c.String(200, "hello world")
    })
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    router.POST("/submit", func(c *gin.Context) {
        c.String(401, "not authorized")
    })
    router.PUT("/error", func(c *gin.Context) {
        c.String(500, "and error hapenned :(")
    })
    router.GET("/books", func(c *gin.Context) {
        // c.String(200, "and error hapenned :(")
        c.String(200,ShowBooks(db));
    })
    router.Run(":8080")
}
func ShowBooks(db *sql.DB) string {    
        var title, author string
        err := db.QueryRow("select title, author from books").Scan(&title, &author)
        if err != nil {
            panic(err)
        }
        return fmt.Sprintf("The first book is '%s' by '%s'", title, author);
}

func NewDB() *sql.DB {
    db, err := sql.Open("sqlite3", "example.sqlite")
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("create table if not exists books(title text, author text)")
    if err != nil {
        panic(err)
    }

    return db
}
