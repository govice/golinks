package cmd

import (
	"log"

	"github.com/laughingcabbage/golinks/types/blockmap"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/urfave/cli"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link an archive ",
	// Long:  "Build out an archive to test on based according to the preferred size",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(args)
		if err := link(args[0], cmd); err != nil {
			log.Println(err)
			cmd.Help()
		}
	},
}

func link(path string, cmd *cobra.Command) error {
	log.Println("verifying link path")
	if valid, err := verifyPath(path); !valid || (err != nil) {
		if err != nil {
			return err
		}
		return errors.New("link: invalid path to link")
	}

	blkmap := blockmap.New(path)
	log.Println("generating link in " + path)
	if err := blkmap.Generate(); err != nil {
		return cli.NewExitError(err, 0)
	}
	blkmap.PrintBlockMap()

	log.Println("saving blockmap to .link file")
	if err := blkmap.Save(path); err != nil {
		return err
	}
	return nil
}
