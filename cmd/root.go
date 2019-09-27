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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile, userLicense string
	verbose              bool
	rootCmd              = &cobra.Command{
		Use:   "golinks",
		Short: "golinks is a tool used to retain and reord deatiled integrity of an archive",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is the root command")
		},
	}
)

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golinks)")

	rootCmd.AddCommand(buildTestCmd)

	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&setRemoteURL, "remote", "r", "", "URL for the remote API")
	configCmd.Flags().StringVarP(&setTmpPath, "temp", "t", "", "Path to the temporary folder for link storage")
	configCmd.Flags().StringVarP(&setTmpPath, "test", "", "", "Path to the temporary folder for test file generation")
	configCmd.Flags().StringVarP(&setStagingPath, "stage", "", "", "Path to directory used for staging changes")
	configCmd.Flags().StringVarP(&setAuthURL, "auth", "", "", "URL used to verify user authentication")
	configCmd.AddCommand(printConfigCmd)

	buildTestCmd.AddCommand(cleanTestCmd)
	buildTestCmd.Flags().StringVarP(&testSize, "size", "s", "", "Test size [small, medium, large] (required)")
	buildTestCmd.Flags().BoolVarP(&randomize, "randomize", "r", false, "randomize generated test data")
	if err := buildTestCmd.MarkFlagRequired("size"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(walkCmd)

	rootCmd.AddCommand(stageCmd)

	rootCmd.AddCommand(pushCmd)

	rootCmd.AddCommand(statusCmd)

	linkCmd.Flags().BoolVarP(&zipArchive, "zip", "z", false, "zip archive after linking")
	rootCmd.AddCommand(linkCmd)

	rootCmd.AddCommand(validateCmd)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

	authCmd.Flags().StringVarP(&setAuthEmail, "email", "e", "", "Set authentication email")
	authCmd.Flags().StringVarP(&setAuthToken, "token", "t", "", "Set API token")
	authCmd.Flags().BoolVarP(&skipVerification, "skip-validation", "", false, "Skip authentication validation step")
	rootCmd.AddCommand(authCmd)

	rootCmd.AddCommand(pushCmd)

}

func initConfig() {
	// read config from config or from home directory
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	golinksHomeFolder := home + string(os.PathSeparator) + ".golinks"
	os.Mkdir(golinksHomeFolder, 0755)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.

		viper.AddConfigPath(golinksHomeFolder)

	}

	viper.SetConfigType("json")
	viper.SetConfigName("golinks")
	viper.AutomaticEnv()

	verb("Reading Config file")
	if err := viper.ReadInConfig(); err != nil {
		verb("Failed to find config file")
		err := SetDefaultConfig()
		if err != nil {
			log.Fatal("Failed to generate default config", err)
			os.Exit(1)
		}

		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		verb("created default config at " + viper.Get(cConfigPath).(string))
	}
	verb("Using config file:", viper.ConfigFileUsed())
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// verb wraps log and only prints when verbose is enabled
func verb(msg ...interface{}) {
	if verbose {
		log.Println(msg...)
	}
}

// var userLicense = `Copyright 2018-2019 Kevin Gentile

//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at

//  http://www.apache.org/licenses/LICENSE-2.0

//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//  `
