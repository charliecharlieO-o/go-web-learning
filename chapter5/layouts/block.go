package main

import (
  "net/http"
  "html/template"
  "time"
  "math/rand"
  "os"
)

func process(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  rand.Seed(time.Now().Unix())
  var t *template.Template
  if rand.Intn(10) > 5 {
    t, _ = template.ParseFiles(wd + "/block.html", wd + "/red_hello.html")
  } else {
    t, _ = template.ParseFiles(wd + "/block.html")
  }
  t.ExecuteTemplate(w, "layout", "")
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
