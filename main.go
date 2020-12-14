package main
import (
  "github.com/gorilla/mux"
  "database/sql"
  "net/http"
  "fmt"
  "github.com/joho/godotenv"
  "os"
  "log"
  _"github.com/go-sql-driver/mysql"
)
type Post struct {
  USERNAME string `json:"username"`
}
var db *sql.DB
var err error
func init() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

}
func main() {
  DB_HOST := os.Getenv("DB_HOST")
  DB_PORT := os.Getenv("DB_PORT")
  DB_USER := os.Getenv("DB_USER")
  DB_PASSWORD := os.Getenv("DB_PASSWORD")
  DB_NAME := os.Getenv("DB_NAME")
  db, err = sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
  if err != nil {
    panic(err.Error())
  }
  router := mux.NewRouter()
  router.HandleFunc("/user/{id}", getPost).Methods("GET")
  http.ListenAndServe(":80", router)
  defer db.Close()
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
}
