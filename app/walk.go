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
	for _, flagName := range c.FlagNames() {
		if !c.IsSet(flagName) {
			continue
		}
		switch flagName {
		case "print":
			w.PrintArchive()
		case "link":
			if err := generateLink(&w); err != nil {
				return cli.NewExitError("Failed to generate link", 0)
			}
		default:
			w.PrintArchive()
		}
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

//generateLink creates a stores linked files in the provided directory. It does not verify the path.
func generateLink(w *walker.Walker) error {
	fmt.Println("generating link in " + w.Root())

	for _, path := range w.Archive() {
		if rel, err := filepath.Rel(w.Root(), path); err == nil {
			fmt.Println(rel)
		}
		//fmt.Println(path)
	}

	return nil
}
