package main

import (
  "net/http"
  "encoding/json"
)

type Post struct {
  User string
  Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  post := &Post{
    User:     "Charlie",
    Threads:  []string{"first", "second", "third"},
  }
  json, _ := json.Marshal(post)
  w.Write(json)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/json", jsonExample)
  server.ListenAndServe()
}
