package main

import (
  "net/http"
  "html/template"
  "time"
  "os"
)

func formatDate(t time.Time) string  {
  layout := "2006-01-02"
  return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  funcMap := template.FuncMap{ "fdate": formatDate }
  t := template.New("custom-functions.html").Funcs(funcMap)
  t, _ = t.ParseFiles(wd + "/custom-functions.html")
  t.Execute(w, time.Now())
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
