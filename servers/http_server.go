package main

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
  "sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, r.URL.Path[1:])
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
  mutex.Lock()
  counter++
  fmt.Fprintf(w, strconv.Itoa(counter))
  mutex.Unlock()
}

func sayHi(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi")
}

func main() {
  http.HandleFunc("/", echoString)
  http.HandleFunc("/increment", incrementCounter)
  http.HandleFunc("/hi", sayHi)

  log.Fatal(http.ListenAndServe(":8080", nil))
}
