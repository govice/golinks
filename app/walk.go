package main

import (
	"os"
	"path/filepath"

	"fmt"

	"github.com/LaughingCabbage/goLinks/types/walker"
	"github.com/urfave/cli"
)

func appWalk(c *cli.Context) error {
	//Verify provided walk path is valid
	if err := verifyPath(c.Args().First()); err != nil {
		return err
	}
	//Extract the valid path
	path, err := filepath.Abs(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}
	//Construct new walker and begin execution
	w = walker.New(path)

	if err := w.Walk(); err != nil {
		return cli.NewExitError(err.Error(), 0)
	}

	//Process post walk commands

	if c.Bool("print") {
		fmt.Println("set")
		w.PrintArchive()
	} else {
		w.PrintArchive()
	}
	return nil
}

func verifyPath(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return cli.NewExitError("Add: archive root does not exit", 0)
		}
		panic(err)
	}
	return nil
}
