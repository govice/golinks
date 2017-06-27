package main

import (
	"fmt"

	"github.com/LaughingCabbage/goLinks/network/server"
)

func main() {
	fmt.Println("executing server...")
	server.Listen()
}
