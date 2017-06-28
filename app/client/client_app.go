package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("executing client...")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to connect ", err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("text to send: ")
		data, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, data+"\n")

		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server reply: " + msg)
	}
}
