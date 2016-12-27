package main

import (
  "fmt"
  "net/http"
  //"net/url"
  "io/ioutil"
  "strings"
	"log"
)
func keepLines(s string, n int) string {
  result := strings.Join(strings.Split(s, "\n")[:n], "\n")
  return strings.Replace(result, "\r", "", -1)
}

func main() {
  resp, err := http.Get("http://localhost:8080/calc")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println("get:\n", keepLines(string(body),1))
/*
  resp, err = http.PostForm("http://localhost:8080", url.Values{"q": {"github"}})
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err = ioutil.ReadAll(resp.Body)
  fmt.Println("Post:\n", keepLines(string(body), 3))
  */
}
