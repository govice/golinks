package server

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	go Listen()

	time.Sleep(10)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Error("Network dial failed")
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("status", status)
	conn.Close()
}
