package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "fmt"
)

//Post Type
type Post struct {
		ID 			int  `json:"id"`
    Title   string   `json:"title,omitempty"`
    Body    string   `json:"body,omitempty"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}


var posts []Post
var db *sql.DB

// Display all from the people var
func GetPosts(w http.ResponseWriter, r *http.Request) {
    // params := mux.Vars(r)
    // perform a db.Query insert
    var arr_posts []Post
    rows , err := db.Query("SELECT id, title, body, created_at, updated_at from posts;")
    if(err != nil){
    	log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
			var post Post
			if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
				// Check for a scan error.
				// Query rows will be closed with defer.
				log.Fatal(err)
			}
			arr_posts = append(arr_posts, post)
		}
		json.NewEncoder(w).Encode(arr_posts)
}

// // Display a single data
// func GetPerson(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     for _, item := range people {
//         if item.ID == params["id"] {
//             json.NewEncoder(w).Encode(item)
//             return
//         }
//     }
//     json.NewEncoder(w).Encode(&Person{})
// }

// create a new item
func CreatePost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var query string
    fmt.Println(params)
    // var body_post Post
    // _ = json.NewDecoder(r.Body).Decode(&body_post)
    // perform a db.Query insert
    query = fmt.Sprintf(
    	"INSERT INTO posts (title, body, created_at, updated_at) VALUES('%s', '%s', NOW(), NOW())",
    	 r.FormValue("title"), r.FormValue("body"))
    _ , err := db.Query(query)
    if(err != nil){
    	log.Fatal(err)
    } else {
    	w.WriteHeader(http.StatusOK)
    }
}

// Delete an item
func DeletePost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var query string
    query = fmt.Sprintf("DELETE from posts where id = %s", params["id"])
    _, err := db.Query(query)
    if(err != nil){
    	log.Fatal(err)
    } else {
    	w.WriteHeader(http.StatusOK)	
    }

}

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