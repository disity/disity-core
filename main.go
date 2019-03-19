package main

import (
    "github.com/gorilla/mux"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
)

var db *sql.DB

// main function to boot up everything
func main() {
   // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    var err error
    db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/disity")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
    router := mux.NewRouter()
    router.HandleFunc("/posts", GetPosts).Methods("GET")
    // router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/posts", CreatePost).Methods("POST")
    router.HandleFunc("/posts/{id}", DeletePost).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}