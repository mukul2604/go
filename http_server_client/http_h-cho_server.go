package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
	"os"
	"encoding/csv"
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


func LookupClient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s := r.FormValue("cname")
	client := strings.TrimSpace(s)
	// package os open. same for all files. returns *File  associated descriptor
	// it is in only read mode use OpenFile for create/write and other ops
	f, err := os.Open("/Users/mukul/work/src/http_server_client/order_status.txt")
	if err != nil {
		panic(err)
	}
	defer  f.Close()

	//csv reader
	csvrdr := csv.NewReader(f)
	rows, err := csvrdr.ReadAll()

	if err != nil {
		panic(err)
	}

	//dumb lookup for the the client's order details
	fmt.Fprintf(w, "Looking into the database...\n")
	//adding artificial lookup cost
	time.Sleep(500 * time.Millisecond)

	for _, row := range rows {
		if strings.Compare(client, row[0]) == 0 {
			fmt.Println("Client Found",client)
			fmt.Fprintf(w, "Here is your Order Status..\n")
			fmt.Fprintf(w, "Client Details: %s %s %s\nOrder Status: %s",
				row[0],	row[1],
				row[2], row[3])
			return
		}
	}
	fmt.Fprintf(w, "Not found in our database")
}

func main() {
	http.HandleFunc("/", EchoWelcome)
	http.HandleFunc("/palindrome", IsPalindrome)
	http.HandleFunc("/date", DateAndTime)
	http.HandleFunc("/clientr",LookupClient)
  	fmt.Println("Server started...")
	//Listening on localhost:8080
	// nil means using DefaultServeMux
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
