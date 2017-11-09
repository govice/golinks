package server

import (
	"log"
	"net"
)

func Listen() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic("Listener binding to tcp failed")
	}

	//for {
	connection, err := listener.Accept()
	if err != nil {
		log.Panic("listner returned error attmepting to accept connection")
	}

	//}
	connection.Close()
}
