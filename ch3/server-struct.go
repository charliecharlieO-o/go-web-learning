package main

import (
  "fmt"
  "net/http"
)

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
  handler := MyHandler{}
  server := http.Server{
    Addr:       "127.0.0.1:8000",
    Handler:    &handler,
  }
  fmt.Println("Server listening")
  server.ListenAndServe()
}
