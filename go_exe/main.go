package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// var posts []Post
var db *sql.DB
var err error
var host = "http://localhost"
var port = "12345"
var connectionString = "root:Welcome#$123@tcp(127.0.0.1:3306)/divya"

// db, _ := sql.Open("mysql", "root:Welcome#$123@tcp(127.0.0.1:3306)/divya")

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := db.Query("SELECT * from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)

}

func createPost(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var post Post
	// _ = json.NewDecoder(r.Body).Decode(&post)
	// // post.ID = strconv.Itoa(rand.Intn(1000000))
	// posts = append(posts, post)
	// json.NewEncoder(w).Encode(&post)
	stmt, err := db.Prepare("INSERT INTO posts(id, title, body) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	id := keyVal["id"]
	title := keyVal["title"]
	body := keyVal["body"]
	_, err = stmt.Exec(id, title, body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// params := r.URL.Query().Get("id")
	// for _, item := range posts {
	// 	if item.ID == params {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Post{})
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query().Get("id")
	result, err := db.Query("SELECT * FROM posts WHERE id = ?", params)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// for index, item := range posts {
	// 	if item.ID == params["id"] {
	// 		posts = append(posts[:index], posts[index+1:]...)
	// 		var post Post
	// 		_ = json.NewDecoder(r.Body).Decode(&post)
	// 		post.ID = params["id"]
	// 		posts = append(posts, post)
	// 		json.NewEncoder(w).Encode(&post)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(posts)
	params := r.URL.Query().Get("id")
	stmt, err := db.Prepare("UPDATE posts SET title = ?, body = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]
	newBody := keyVal["body"]
	_, err = stmt.Exec(newTitle, newBody, params)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params)

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// for index, item := range posts {
	// 	if item.ID == params["id"] {
	// 		posts = append(posts[:index], posts[index+1:]...)
	// 		break
	// 	}
	// }
	// json.NewEncoder(w).Encode(posts)
	params := r.URL.Query().Get("id")
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was deleted", params)
}
func main() {

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts(id INT NOT NULL, title VARCHAR(20), body VARCHAR(255), PRIMARY KEY (ID));")
	defer db.Close()
	router := mux.NewRouter()

	// posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	// posts = append(posts, Post{ID: "2", Title: "My second post", Body: "This is the content of my second post"})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/id", getPost).Methods("GET")
	router.HandleFunc("/posts/update", updatePost).Methods("PUT")
	router.HandleFunc("/posts/delete", deletePost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
