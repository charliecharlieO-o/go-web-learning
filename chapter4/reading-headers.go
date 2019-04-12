// Basic header methods: add, delete, get and set
// Request headrs are a map of strings as keys and slices/arrays of strings as values
package main

import (
  "fmt"
  "net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
  h := r.Header
  // To print a specific header
  t := r.Header["Accept-Language"]
  fmt.Fprintln(w, h, t)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/headers", headers)
  server.ListenAndServe()
}
