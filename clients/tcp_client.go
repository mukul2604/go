package main

import  "fmt"
import  "net"
import  "bufio"
import  "os"

func main() {
  fmt.Printf("hello, client\n")
  //connect to the socket
  conn, _ := net.Dial("tcp","127.0.0.1:8081")
  for {
    //reader in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    //send to socket
    fmt.Fprintf(conn, text + "\n")
    //listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}
