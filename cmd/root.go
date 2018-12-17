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
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile, userLicense string

	rootCmd = &cobra.Command{
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
	// rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output") // TODO provide verbose output

	rootCmd.AddCommand(buildTestCmd)

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(printConfigCmd)

	buildTestCmd.AddCommand(cleanTestCmd)
	buildTestCmd.Flags().StringVarP(&testSize, "size", "s", "", "Test size [small, medium, large] (required)")
	if err := buildTestCmd.MarkFlagRequired("size"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(walkCmd)

	linkCmd.Flags().BoolVarP(&zipArchive, "zip", "z", false, "zip archive after linking")
	rootCmd.AddCommand(linkCmd)

	rootCmd.AddCommand(validateCmd)
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

	log.Println("Reading Config file")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Failed to find config file")
		config, err := DefaultConfig()
		if err != nil {
			log.Fatal("Failed to generate default config", err)
			os.Exit(1)
		}
		log.Println("Creating Default Config at " + config.ConfigPath)
		if err := config.WriteConfig(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	log.Println("Using config file:", viper.ConfigFileUsed())
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// var userLicense = `Copyright 2018-2018 Kevin Gentile

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
