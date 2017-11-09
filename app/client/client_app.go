package main

import (
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
	app := cli.NewApp()
	app.Name = "goLinks Client"
	app.Usage = "usage"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for greeting",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := "Kev"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
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
