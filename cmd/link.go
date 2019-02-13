/*
 *Copyright 2018-2019 Kevin Gentile
 *
 *Licensed under the Apache License, Version 2.0 (the "License");
 *you may not use this file except in compliance with the License.
 *You may obtain a copy of the License at
 *
 *http://www.apache.org/licenses/LICENSE-2.0
 *
 *Unless required by applicable law or agreed to in writing, software
 *distributed under the License is distributed on an "AS IS" BASIS,
 *WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *See the License for the specific language governing permissions and
 *limitations under the License.
 */

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/govice/golinks/blockmap"
	"github.com/pierrre/archivefile/zip"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/urfave/cli"
)

var zipArchive bool

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link an archive ",
	// Long:  "Build out an archive to test on based according to the preferred size",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(args)
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
			fmt.Println("ZIP ARCHIVE")
			if err := zipArchiveF(archivePath); err != nil {
				log.Println(err)
				// cmd.Help()
			}
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
