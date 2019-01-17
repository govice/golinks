/*
 *Copyright 2018 Kevin Gentile
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

	"github.com/laughingcabbage/golinks/walker"
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
