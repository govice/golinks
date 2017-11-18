package main

import (
	"os"

	"github.com/LaughingCabbage/goLinks/types/walker"
	"github.com/urfave/cli"
)

const (
	//walkerWorkers int    = 1
	defaultRoot string = ""
	//configName  string = "config"
)

var outputDir string = defaultRoot

var w = walker.New(defaultRoot)

func main() {
	app := cli.NewApp()
	app.Name = "golinks"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		{
			Name:  "Kevin Gentile",
			Email: "kevin@kevingentile.com",
		},
	}
	app.Copyright = "(c) 2017 Kevin Gentile"
	app.HelpName = "golinks"
	app.Usage = "a blockchain for your filesystem"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "set output 'DIR'",
			Destination: &outputDir,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "walk",
			Aliases:     []string{"w"},
			Usage:       "walk a given archive",
			Description: "walk a given archive. print by default",
			Action:      appWalk,
		},
	}

	//legooo
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
