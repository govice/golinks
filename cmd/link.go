package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/govice/golinks/blockmap"
	"github.com/pierrre/archivefile/zip"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var zipArchive bool

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link an archive ",
	// Long:  "Build out an archive to test on based according to the preferred size",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("link: missing arguments for command")
			cmd.Help()
		}
		archivePath := args[0]
		if err := link(archivePath, cmd); err != nil {
			log.Println(err)
			cmd.Help()
		}

		if zipArchive {
			verb("zipping archive: " + archivePath)
			if err := zipArchiveF(archivePath); err != nil {
				log.Println(err)
			}
		}
	},
}

func link(path string, cmd *cobra.Command) error {
	verb("verifying link path")
	if valid, err := verifyPath(path); !valid || (err != nil) {
		if err != nil {
			return err
		}
		return errors.New("link: invalid path to link")
	}

	blkmap := blockmap.New(path)
	verb("generating link in " + path)
	if err := blkmap.Generate(); err != nil {
		return cli.NewExitError(err, 0)
	}

	if verbose {
		blkmap.PrintBlockMap()
	}

	verb("saving blockmap to .link file")
	rootHash := base64.StdEncoding.EncodeToString(blkmap.RootHash)
	verb("Root hash: " + rootHash)
	tmpLinkPath := viper.Get(cTempPath).(string)
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	if err := blkmap.SaveNamed(tmpLinkPath, uuid.String()); err != nil {
		return err
	}
	return nil
}

func zipArchiveF(path string) error {
	archive, err := os.Open(path)
	if err != nil {
		return err
	}

	defer archive.Close()
	outPath := path + ".zip"
	progress := func(outPath string) {
		fmt.Println(outPath)
	}

	err = zip.ArchiveFile(path, outPath, progress)
	if err != nil {
		return err
	}

	return nil
}
