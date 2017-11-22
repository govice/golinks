package main

import (
	"os"

	"github.com/LaughingCabbage/goLinks/types/walker"
	"github.com/urfave/cli"
)

const (
	//walkerWorkers int    = 1
	defaultRoot string = "r"
	//configName  string = "config"
)

var outputDir = defaultRoot

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

	app.Commands = []cli.Command{
		{
			Name:        "walk",
			Aliases:     []string{"w"},
			Usage:       "walk a given archive",
			Description: "walk a given archive. print by default",
			Action:      appWalk,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "output, o",
					Usage:       "set output 'DIR'",
					Destination: &outputDir,
				},
				cli.BoolFlag{
					Name:  "print, p",
					Usage: "print walked archive",
				},
				cli.BoolFlag{
					Name:  "link, l",
					Usage: "generate link for walked directory",
				},
			},
		},
		{
			Name:        "maketest",
			Aliases:     []string{"bt"},
			Usage:       "build test directory",
			Description: "builds a test directory in the provided location",
			Action:      appBuildTestDir,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "small, s",
					Usage: "small directory",
				},
				cli.BoolFlag{
					Name:  "medium, m",
					Usage: "medium directory",
				},
				cli.BoolFlag{
					Name:  "large, l",
					Usage: "large directory",
				},
			},
		},
	}

	//legooo
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}