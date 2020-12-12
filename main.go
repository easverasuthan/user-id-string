package main
import (
  "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "fmt"
)
type Post struct {
  USERNAME string `json:"username"`
}
var db *sql.DB
var err error
func main() {
  db, err = sql.Open("mysql", "admin:8iIXKcaC0zc3UuVb5gSu@tcp(pratilipi.ctnpjkz144vl.us-east-1.rds.amazonaws.com:3306)/pratilipi")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  router := mux.NewRouter()
  router.HandleFunc("/users/{id}", getPost).Methods("GET")
  http.ListenAndServe(":80", router)
}
func getPost(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  result, err := db.Query("SELECT username FROM users WHERE user_id = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  var post Post
  for result.Next() {
    err := result.Scan(&post.USERNAME)
    if err != nil {
      panic(err.Error())
    }
  }
  fmt.Fprintf(w, post.USERNAME)
