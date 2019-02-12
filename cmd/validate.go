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
	log.Println("verifying link path")
	if valid, err := verifyPath(path); !valid || (err != nil) {
		if err != nil {
			return err
		}
		return errors.New("link: invalid path to link")
	}

	//Load blockmap from existing file
	log.Println("checking for existing link file")
	fileBlockmap := blockmap.New(path)
	if err := fileBlockmap.Load(path); err != nil {
		return err
	}

	//Validate the existing directory
	log.Println("validating link file with current archive")
	temp := blockmap.New(path)
	if err := temp.Generate(); err != nil {
		return err
	}

	//Compare file with existing directory
	if !blockmap.Equal(fileBlockmap, temp) {
		return errors.New("invalid link")
	}
	log.Println("link is valid")
	return nil
}
