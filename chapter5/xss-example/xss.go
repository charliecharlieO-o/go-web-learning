package main

import (
  "net/http"
  "html/template"
  "os"
)

func process(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  t, _ := template.ParseFiles(wd + "/tmpl.html")
  //  t.Execute(w, r.FormValue("comment")) standar
  // Bypassing Go's template context
  t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  t, _ := template.ParseFiles(wd + "/xss-form.html")
  t.Execute(w, nil)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  http.HandleFunc("/form", form)
  server.ListenAndServe()
}
