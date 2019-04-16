/*
  Go's templating language is context-aware, that means the way it's displayed changes
  depending on its location in the document, this to help scape special characters in
  things like HTML, CSS and JavaScript and that way avoiding basic injections.
*/
package main

import (
  "net/http"
  "html/template"
  "os"
)

func process(w http.ResponseWriter, r *http.Request) {
  wd, _ := os.Getwd()
  t, _ := template.ParseFiles(wd + "/context.html")
  content := `I asked: <i>"What's up?"</i>`
  t.Execute(w, content)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
