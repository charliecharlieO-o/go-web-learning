package main

import (
  "net/http"
  "html/template"
  "os"
)

func process(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  t, _ := template.ParseFiles(wd + "/template-intro/tmpl.html")
  t.Execute(w, "Hello World!")
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
