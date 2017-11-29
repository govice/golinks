package main

import (
	"os"

	"path/filepath"

	"github.com/urfave/cli"
)

func verifyPath(path string) (string, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return "", cli.NewExitError("Add: archive root does not exist", 0)
		}
		panic(err)
	}

	//Extract the valid path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}
