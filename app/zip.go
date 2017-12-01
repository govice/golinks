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
		return cli.NewExitError("invalid source path", 0)
	}
	targetPath, err := verifyPath(c.Args().Get(1))
	if err != nil {
		return cli.NewExitError("invalid target path", 0)
	}

	log.Println("compressing file(s)")
	if err := fs.Compress(dirPath, targetPath); err != nil {
		return cli.NewExitError(err, 0)
	}

	log.Println("compressed archive to " + targetPath)
	return nil
}

func appUnzip(c *cli.Context) error {
	log.Println("verifying directories")
	archivePath, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError("invalid source path", 0)
	}
	targetPath, err := verifyPath(c.Args().Get(1))
	if err != nil {
		return cli.NewExitError("invalid target path", 0)
	}

	log.Println("decompressing file(s)")
	if err := fs.Decompress(archivePath, targetPath); err != nil {
		return cli.NewExitError("failed to decompress files", 0)
	}
	log.Println("decompressed archive to " + targetPath)

	return nil
}
