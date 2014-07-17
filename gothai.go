package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
  "io/ioutil"
)

type Message struct {
  Word string
}

func main() {
  http.HandleFunc("/submitWord", handler)
  fmt.Println("Listening...")
  err := http.ListenAndServe(GetPort(),nil)
  if err != nil {
    log.Fatal("ListenAndServe: ",  err)
    return
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
  w.Header().Set("Access-Control-Allow-Credentials", "true")

  body, err := ioutil.ReadAll(r.Body)
  var msg Message
  err = json.Unmarshal(body, &msg)
  _ = err
  fmt.Println(msg)
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

