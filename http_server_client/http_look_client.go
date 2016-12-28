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
	//defer keyword is used to defer the close
	// it will be definitely called at the end.
  	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	for {
		//Input reading from the client
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the Client name: ")
		cname, _ := reader.ReadString('\n')

		// Post the expression evaluation request
		// url.Values is used to pass the parameter,val
		resp, err = http.PostForm("http://localhost:8080/clientr",
			url.Values{"cname": {string(cname)}})
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}
