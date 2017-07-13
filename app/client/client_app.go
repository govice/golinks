package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

//var gopherType string

const (
	defaultGopher = "pocket"
	usage         = "the variety of gopher"
)

func main() {
	//gopherType := flag.String("command", "cmd", "This is the help menu")

	goph := flag.String("word", "bar", "a string")

	//flag.StringVar(gopherType, "gopher_type", defaultGopher, usage)
	//flag.StringVar(gopherType, "g", defaultGopher, usage+" (shorthand)")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*goph)
	//fmt.Println(stringFlag)
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	app.Run(os.Args)
}

/*
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
*/
