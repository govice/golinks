package main

import (
	"github.com/LaughingCabbage/goLinks/network/client"
)

func main() {
	peer := client.New("localhost", "8080")
	err := peer.Connect()
	if err != nil {
		panic(err)
	}

	for {
		peer.Message("TEST MESSAGE")
		peer.Listen()
	}
}
