package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
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

func GetPosts(w http.ResponseWriter, r *http.Request) {
    var arr_posts []Post
    rows , err := db.Query("SELECT id, title, body, created_at, updated_at from posts;")
    if(err != nil){
    	log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
			var post Post
			if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
				log.Fatal(err)
			}
			arr_posts = append(arr_posts, post)
		}
		json.NewEncoder(w).Encode(arr_posts)
}


func CreatePost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var query string
    fmt.Println(params)
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