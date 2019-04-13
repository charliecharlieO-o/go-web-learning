package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(1024)
  // Check if data was sent on multipart enc
  if r.MultipartForm == nil {
    fmt.Fprintln(w, "No data sent")
    return
  }
  if len(r.MultipartForm.File) == 0 {
    fmt.Fprintln(w, "No file")
    return
  }
  fileHeader := r.MultipartForm.File["uploaded"][0]
  file, err := fileHeader.Open()
  if err == nil {
    data, err := ioutil.ReadAll(file)
    if err == nil {
      fmt.Fprintln(w, string(data))
    }
  }
}

func main() {
  server := http.Server{
    Addr:     "127.0.0.1:8000",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}
