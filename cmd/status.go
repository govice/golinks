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
