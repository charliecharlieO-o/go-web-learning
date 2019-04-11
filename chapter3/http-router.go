// Remember to checkout the gorilla web toolkit

package main

import (
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintf(w, "Hello, %s!\n", p.ByName("name"))
}

func main() {
  mux := httprouter.New()
  mux.GET("/hello/:name", hello)

  server := http.Server{
    Addr:     "127.0.0.1:8000",
    Handler:  mux,
  }
  server.ListenAndServe()
}
