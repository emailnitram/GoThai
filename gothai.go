package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Listening...")
  err := http.ListenAndServe(GetPort(),nil)
  if err != nil {
    log.Fatal("ListenAndServe: ",  err)
    return
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello this is our first Go web program!")
}

func GetPort() string {
  var port = os.Getenv("PORT")
  if port == "" {
    port = "4747"
    fmt.Println("INFO: No Port environment variable detected, defaulting to " + port)
  }
  return ":" + port
}

