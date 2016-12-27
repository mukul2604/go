package main

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
  "sync"
  "encoding/json"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
  mutex.Lock()
  counter++
  fmt.Fprintf(w, strconv.Itoa(counter))
  mutex.Unlock()
  fmt.Fprintf(w, "Hi")
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  myItems := []string{"item1", "item2", "item3"}
  a,_ := json.Marshal(myItems)
  w.Write(a)
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

func handlerjson(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  myItems := []string{"item1", "item2", "item3"}
  a,_ := json.Marshal(myItems)
  w.Write(a)
}

func main() {
  http.HandleFunc("/", echoString)
  //http.HandleFunc("/increment", incrementCounter)
  //http.HandleFunc("/hi", sayHi)
  //http.HandleFunc("/son", handlerjson)

  log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
