package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "net/http"	
    "fmt"
)

//Comment Type
type Comment struct {
	ID 			int  `json:"id"`
    UserId   int   `json:"user_id"`
    ParentId   int   `json:"parent_id"`
    ParentType string `json:"parent_type`
    Body    string   `json:"body"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

var comments []Comment

func GetComments(w http.ResponseWriter, r *http.Request) {
    var arr_comments []Comment
    rows , err := db.Query("SELECT id, user_id, parent_id,parent_type,body, created_at, updated_at from comments;")
    if(err != nil){
    	log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
			var comment Comment
			if err := rows.Scan(&comment.ID, &comment.UserId, &comment.ParentId, &comment.ParentType, &comment.Body, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
				log.Fatal(err)
			}
			arr_comments = append(arr_comments, comment)
		}
		json.NewEncoder(w).Encode(arr_comments)
}


func CreateComment(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var query string
    fmt.Println(params)
    query = fmt.Sprintf(
    	"INSERT INTO comments (body, user_id, parent_id, parent_type, created_at, updated_at) VALUES('%s', '%s', '%s', '%s', NOW(), NOW())",
    	r.FormValue("body"), r.FormValue("user_id"), r.FormValue("parent_id"), r.FormValue("parent_type"))
    _ , err := db.Query(query)
    if(err != nil){
    	log.Fatal(err)
    } else {
    	w.WriteHeader(http.StatusOK)
    }
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var query string
    query = fmt.Sprintf("DELETE from comments where id = %s", params["id"])
    _, err := db.Query(query)
    if(err != nil){
    	log.Fatal(err)
    } else {
    	w.WriteHeader(http.StatusOK)	
    }

}