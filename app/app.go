package main

import (
	"os"

	"fmt"

	"path/filepath"

	"github.com/LaughingCabbage/goLinks/types/walker"
	"github.com/urfave/cli"
)

const (
	//walkerWorkers int    = 1
	defaultRoot string = ""
	configName  string = "config"
)

var w = walker.New(defaultRoot)

func main() {
	app := cli.NewApp()
	app.Name = "golink"
	app.Usage = "filesystem blockchain"

	app.Commands = []cli.Command{
		{
			Name:    "root",
			Aliases: []string{"r"},
			Usage:   "set walker archive root",
			Action:  appSetRoot,
		},
		{
			Name:    "walk",
			Aliases: []string{"w"},
			Usage:   "walk and print archive",
			Action:  appWalk,
		},
	}

	//legooo
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}

func appSetRoot(c *cli.Context) error {
	fmt.Printf("Adding walker root: %v\n", c.Args().First())
	//Validate archive root
	if _, err := os.Stat(c.Args().First()); err != nil {
		if os.IsNotExist(err) {
			return cli.NewExitError("Add: archive root does not exit", 0)
		}
		panic(err)
	}

	//TODO is this check needed or can c.args be used?
	//Extract valid path for filesystem.
	path, err := filepath.Abs(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}
	w = walker.New(path)
	return nil
}

func appWalk(c *cli.Context) error {
	if err := w.Walk(); err != nil {
		return cli.NewExitError(err.Error(), 0)
	}
	w.PrintArchive()
	return nil
}
