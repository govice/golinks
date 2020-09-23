package cmd

import (
	"fmt"
	"log"

	"github.com/govice/golinks/blockmap"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a linked archive",
	// Long:  "Build out an archive to test on based according to the preferred size",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validate(args[0], cmd); err != nil {
			log.Println(err)
			cmd.Help()
		}
	},
}

//TODO is this re-creating an existing link file?
func validate(path string, cmd *cobra.Command) error {
	//Validate provided path
	verb("verifying link path")
	if valid, err := verifyPath(path); !valid || (err != nil) {
		if err != nil {
			return err
		}
		return errors.New("link: invalid path to link")
	}

	//Load blockmap from existing file
	verb("checking for existing link file")
	fileBlockmap := blockmap.New(path)
	if err := fileBlockmap.Load(path); err != nil {
		return err
	}

	//Validate the existing directory
	verb("validating link file with current archive")
	temp := blockmap.New(path)
	if err := temp.Generate(); err != nil {
		return err
	}

	//Compare file with existing directory
	if !blockmap.Equal(fileBlockmap, temp) {
		return errors.New("invalid link")
	}

	fmt.Println("link is valid")
	return nil
}
