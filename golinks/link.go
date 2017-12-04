package main

import (
	"log"

	"github.com/LaughingCabbage/goLinks/types/blockmap"
	"github.com/urfave/cli"
)

func appLink(c *cli.Context) error {
	log.Println("running linker")
	absPath, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}

	b := blockmap.New(absPath)
	log.Println("generating link in " + absPath)
	if err := b.Generate(); err != nil {
		return cli.NewExitError(err, 0)
	}
	b.PrintBlockMap()

	log.Println("saving blockmap to .link file")
	if err := b.Save(absPath); err != nil {
		return cli.NewExitError(err, 0)
	}
	return nil
}
