package main

import (
	"log"

	"github.com/LaughingCabbage/goLinks/types/blockmap"
	"github.com/urfave/cli"
)

func appValidate(c *cli.Context) error {
	//Validate provided path
	path, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}

	//Load blockmap from existing file
	log.Println("checking for existing link file")
	b := &blockmap.BlockMap{}
	if err := b.Load(path); err != nil {
		return cli.NewExitError("link file not found", 0)
	}

	//Validate the existing directory
	log.Println("validating link file with current archive")
	temp := blockmap.New(path)
	if err := temp.Generate(); err != nil {
		return cli.NewExitError(err, 0)
	}

	//Compare file with existing directory
	if !blockmap.Equal(b, temp) {
		return cli.NewExitError("link is invalid", 0)
	}

	log.Println("link is valid")
	return nil
}
