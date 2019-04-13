package main

import (
  "fmt"
  "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
  // You need to tell the method how much data you want to extract
  // in bytes
  // Note: When using multipart form enc no url params are added

  // r.ParseMultipartForm(1024)
  // fmt.Fprintln(w, r.MultipartForm) Access the Form field

  // FormValue method retreives the first value & you don't need to parse the form first
  // fmt.Fprintln(w, r.FormValue("hello"))
  // fmt.Fprintln(w, r.Form)

  fmt.Fprintln(w, "(1)", r.FormValue("hello"))
  fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
  fmt.Fprintln(w, "(3)", r.PostForm)
  fmt.Fprintln(w, "(4)", r.MultipartForm)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
