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
	"io/ioutil"

	"github.com/urfave/cli"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "status of recently generated link files",
		Run: func(cmd *cobra.Command, args []string) {
			tempPath := viper.Get(cTempPath).(string)
			stagePath := viper.Get(cStagingPath).(string)

			stagedFiles, err := ioutil.ReadDir(stagePath)
			if err != nil {
				cli.NewExitError(err, 1)
			}

			linkFiles, err := ioutil.ReadDir(tempPath)
			if err != nil {
				cli.NewExitError(err, 1)
			}

			if len(stagedFiles) > 0 {
				verb("Staged Files: ", len(stagedFiles))
				fmt.Println("Staged:")
				for _, info := range stagedFiles {
					fmt.Println("  " + info.Name())
				}
			}

			if len(linkFiles) > 0 {
				verb("Linked Files: ", len(linkFiles))
				fmt.Println("Linked:")
				for _, info := range linkFiles {
					fmt.Println("  " + info.Name())
				}
			}
		},
	}
)
