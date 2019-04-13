/*
  If the URL  and the form have the same key the will be both placed in a slice
  with the form value always prioritized before the URL value.
*/

package main

import (
  "fmt"
  "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
  r.ParseForm() // Parse the request form
  fmt.Fprintln(w, r.Form) // Access the Form field
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
