// Go package
package main

/// Go fmt import
import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

// Go main function
func main() {
    fmt.Println("Hello World!")
}

// db config
const (
    DB_USER     = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME     = "movies"
)

// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return DB
}