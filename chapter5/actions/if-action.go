package main

import (
  "net/http"
  "html/template"
  "math/rand"
  "time"
  "os"
)

func process(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  t, _ := template.ParseFiles(wd + "/if-action.html")
  rand.Seed(time.Now().Unix())
  t.Execute(w, rand.Intn(10) > 5)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
