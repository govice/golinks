package main

import (
	"github.com/LaughingCabbage/goLinks/types/walker"
	"github.com/urfave/cli"
)

func appWalk(c *cli.Context) error {
	//Verify provided walk path is valid
	path, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}

	//Construct new walker and begin execution
	w = walker.New(path)

	if err := w.Walk(); err != nil {
		return cli.NewExitError(err.Error(), 0)
	}

	//Process post walk commands
	for _, flagName := range c.FlagNames() {
		if !c.IsSet(flagName) {
			continue
		}
		switch flagName {
		case "print":
			w.PrintArchive()
		default:
			w.PrintArchive()
		}
	}
	return nil
}
