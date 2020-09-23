package cmd

import (
	"os"

	"github.com/urfave/cli"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	stageCmd = &cobra.Command{
		Use:   "stage",
		Short: "stage link file",
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			verb("staging: " + fileName)
			tmpPath := viper.Get(cTempPath).(string)
			stagePath := viper.Get(cStagingPath).(string)
			err := os.Rename(tmpPath+string(os.PathSeparator)+fileName, stagePath+string(os.PathSeparator)+fileName)
			if err != nil {
				cli.NewExitError(err, 1)
			}
		},
	}
)
