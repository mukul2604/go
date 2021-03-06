package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"bufio"
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

	for {
		//Input reading from the client
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the expression: ")
		expr, _ := reader.ReadString('\n')

		// Post the expression evaluation request
		// url.Values is used to pass the parameter,val
		resp, err = http.PostForm("http://localhost:8080/palindrome",
			url.Values{"expr": {string(expr)}})
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}
