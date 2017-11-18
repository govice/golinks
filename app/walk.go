package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/LaughingCabbage/goLinks/types/walker"
	"github.com/urfave/cli"
)

func appWalk(c *cli.Context) error {
	if err := setRoot(c); err != nil {
		return err
	}
	if err := w.Walk(); err != nil {
		return cli.NewExitError(err.Error(), 0)
	}
	w.PrintArchive()
	return nil
}

func setRoot(c *cli.Context) error {
	fmt.Printf("Adding walker root: %v\n", c.Args().First())
	//Validate archive root
	if err := verifyPath(c.Args().First()); err != nil {
		return err
	}

	//TODO is this check needed or can c.args be used?
	//Extract valid path for filesystem.
	path, err := filepath.Abs(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}
	w = walker.New(path)
	fmt.Println(c.FlagNames())
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

func appOutput(c *cli.Context) error {
	fmt.Println("checkout output")
	return nil
}
