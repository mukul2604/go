package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
)



func EchoWelcome (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to H-CHO Server..")
}


func IsPalindrome(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	s := r.FormValue("expr")
	expr := strings.TrimSpace(s)
	length := len(expr)
	fmt.Println(length ,expr)
	for i:=0; i < length/2; i++ {
		if expr[i] != expr[length - i - 1] {
			fmt.Println(length, expr[i], expr[length-i-1])
			fmt.Fprintf(w, "Expression %s is not palindrome.", expr)
			return
		}
	}
	fmt.Fprintf(w, "Expression %s is a palindrome.", expr)
}


func DateAndTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "Server Date and Time is %s .", t.Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/", EchoWelcome)
	http.HandleFunc("/palindrome", IsPalindrome)
	http.HandleFunc("/date", DateAndTime)
  fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
