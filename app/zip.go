package main

import (
	"log"

	"github.com/LaughingCabbage/goLinks/types/fs"
	"github.com/urfave/cli"
)

func appZip(c *cli.Context) error {
	log.Println("verifying directories")
	dirPath, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}
	targetPath, err := verifyPath(c.Args().Get(1))
	if err != nil {
		return cli.NewExitError(err, 0)
	}

	log.Println("compressing file(s)")
	if err := fs.Compress(dirPath, targetPath); err != nil {
		return cli.NewExitError(err, 0)
	}

	log.Println("compressed archive to " + targetPath)
	return nil
}
