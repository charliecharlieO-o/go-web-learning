package main

import (
  "fmt"
  "net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
  c1 := http.Cookie{
    Name:     "first_cookie",
    Value:    "Go Web Programming",
    HttpOnly: true,
  }
  c2 := http.Cookie{
    Name:     "second_cookie",
    Value:    "Manning Publications Co",
    HttpOnly: true,
  }
  http.SetCookie(w, &c1)
  http.SetCookie(w, &c2)
}

func getCookieFromHeader(w http.ResponseWriter, r *http.Request) {
  // Gets you a slice with a single string with a cookies formated in
  // if you want one you need to parse it
  h := r.Header["Cookie"]
  fmt.Fprintln(w, h)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
  // Get one cookie
  c1, err := r.Cookie("first_cookie")
  if err != nil {
    fmt.Fprintln(w, "Cannot get the first cookie")
  }
  cs := r.Cookies() // Get all cookies
  fmt.Fprintln(w, c1)
  fmt.Fprintln(w, cs)
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/set-cookie", setCookie)
  http.HandleFunc("/get-cookie", getCookie)
  server.ListenAndServe()
}
