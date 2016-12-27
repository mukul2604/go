package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func EchoWelcome (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to H-CHO Server..")
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	x := r.FormValue("expr")
	fmt.Fprintf(w, "Server is evaluating the expr:%s",x)
	fmt.Println(x)
}

func DateAndTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "Server Date and Time is %s .", t.Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/", EchoWelcome)
  	http.HandleFunc("/calc", Calculator)
	http.HandleFunc("/date", DateAndTime)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
