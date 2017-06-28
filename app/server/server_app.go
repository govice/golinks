package main

import (
	"fmt"

	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("running server...")

	client, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listner could not connect", err)
	}

	conn, err := client.Accept()

	for {
		// will listen for message to process ending in newline (\n)
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(msg))
		// sample process for string received
		newmessage := strings.ToUpper(msg)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
