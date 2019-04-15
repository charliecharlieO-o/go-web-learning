package main

import (
  "net/http"
  "html/template"
  "os"
)

func process(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  t, _ := template.ParseFiles(wd + "/iterator.html")
  daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
  t.Execute(w, daysOfWeek)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
