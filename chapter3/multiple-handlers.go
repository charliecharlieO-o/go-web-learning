package main

import (
  "fmt"
  "net/http"
)

type HelloHandler struct {}

func (*HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Hello!")
}

type WorldHandler struct {}

func (*WorldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "World!")
}

func main() {
  hello := HelloHandler{}
  world := WorldHandler{}

  server := http.Server{
    Addr: "127.0.0.1:8000",
  }

  http.Handle("/hello", &hello)
  http.Handle("/world", &world)

  server.ListenAndServe()
}
