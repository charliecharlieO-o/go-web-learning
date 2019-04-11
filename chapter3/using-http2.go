// Since we are using GO 1.12 to code this examples we no longer need http2
// Accessing this webserver from a modern brwoser will give you a warning
// REMEMBER to access using https otherwise you'll get a:
//    TLS handshake error
package main

import (
  "fmt"
  "net/http"
)

type MyHandler struct {}

func (h *MyHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World!")
}

func main() {
  handler := MyHandler{}
  server := http.Server{
    Addr:     "127.0.0.1:8000",
    Handler:  &handler,
  }
  // Note we are using TLS
  err := server.ListenAndServeTLS("cert.pem", "key.pem")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Server On")
  }
}
