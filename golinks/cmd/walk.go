package cmd

import (
	"log"

	"github.com/laughingcabbage/golinks/types/walker"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var walkPath string

var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "Walk an archive ",
	// Long:  "Build out an archive to test on based according to the preferred size",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(args)
		if err := walk(args[0], cmd); err != nil {
			log.Println(err)
			cmd.Help()
		}
	},
}

func walk(path string, cmd *cobra.Command) error {
	if valid, err := verifyPath(path); !valid || (err != nil) {
		if err != nil {
			return err
		}
		return errors.New("walk: invalid path to walk")
	}

	//Construct new walker and begin execution
	w := walker.New(path)
	if err := w.Walk(); err != nil {
		return err
	}

	w.PrintArchive()
	return nil
}
