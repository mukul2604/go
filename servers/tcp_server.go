package main

import  "fmt"
import  "net"
import  "bufio"
import  "strings"

func main() {
  fmt.Printf("Launching Server...\n")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8081")

  //accept connection on port
  conn, _ := ln.Accept()

  //forever run loop until Ctlr-C
  for {
    //will listen for message to process enfing in newline
    message,_ := bufio.NewReader(conn).ReadString('\n')
    //output message received
    fmt.Print("Message Received:", string(message))
    //sampele process for string received
    newmessage := strings.ToUpper(message)
    //send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }
}
