package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "golink s"
	app.Usage = "filesystem blockchain"

	app.Commands = []cli.Command{
		{
			Name:    "walk",
			Aliases: []string{"w"},
			Usage:   "walk and print archive",
			Action: func(c *cli.Context) error {
				fmt.Println("Walking")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
