/*
  Body field is a io.ReadCloser interface. In turn the Body field consists of a Reader
  interface and a Closer interface. Notice that you first need to determine how much
  bytes to read, then create a byte array of the content length, and call the Read method
  to read into it.
*/

package main

import (
  "fmt"
  "net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
  len := r.ContentLength
  body := make([]byte, len)
  r.Body.Read(body)
  // fmt.Println(string(body)) will print the body string
  // fmt.Println(body) wil print a byte slice
  fmt.Fprintln(w, string(body))
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/body", body)
  server.ListenAndServe()
}
