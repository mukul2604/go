package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	//get the home page
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//get the date and time
  	resp, err = http.Get("http://localhost:8080/date")
  	if err != nil {
    		log.Fatal(err)
  	}
  	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
