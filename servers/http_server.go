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
  	fmt.Fprintf(w, "Hi")
}

func IncrementCounter(w http.ResponseWriter, r *http.Request) {
  	mutex.Lock()
  	counter++
  	fmt.Fprintf(w, strconv.Itoa(counter))
  	mutex.Unlock()
}

func sayHi(w http.ResponseWriter, r *http.Request) {
  	fmt.Fprintf(w, "Hi")
}

func handler2(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json; charset=utf-8")
  	myItems := []string{"item1", "item2", "item3"}
  	a,_ := json.Marshal(myItems)
  	w.Write(a)
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "calling calculator function")
}
func main() {
  	//http.HandleFunc("/", echoString)
	http.HandleFunc("/calc", Calculator)
	http.HandleFunc("/increment", IncrementCounter)
	http.HandleFunc("/hi", sayHi)
	http.HandleFunc("/son", handler2)

  	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
