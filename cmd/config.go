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
	"os/user"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cConfigPath  = "configpath"
	cTestPath    = "testpath"
	cTempPath    = "temporary-folder"
	cStagingPath = "staging-path"
	cRemote      = "remote"
	cAuth        = "auth"
)

var (
	setRemoteURL   string
	setTmpPath     string
	setTestPath    string
	setStagingPath string
	setAuthURL     string
	configCmd      = &cobra.Command{
		Use:   "config",
		Short: "handle tasks related to configuration",
		Run: func(cmd *cobra.Command, args []string) {
			configPath := viper.Get(cConfigPath).(string)
			verb("config path: " + configPath)

			write := false
			if setRemoteURL != "" {
				viper.Set(cRemote, setRemoteURL)
				write = true
			}

			if setTmpPath != "" {
				viper.Set(cTempPath, setTmpPath)
				write = true
			}

			if setTestPath != "" {
				viper.Set(cTestPath, setTestPath)
				write = true
			}

			if setStagingPath != "" {
				viper.Set(cStagingPath, setStagingPath)
				write = true
			}

			if setAuthURL != "" {
				viper.Set(cAuth, setAuthURL)
				write = true
			}

			if write {
				if err := viper.WriteConfig(); err != nil {
					log.Fatal(err)
				}
			} else {
				cmd.Help()
			}
		},
	}

	printConfigCmd = &cobra.Command{
		Use:   "print",
		Short: "print config file",
		Run: func(cmd *cobra.Command, args []string) {
			keys := viper.AllKeys()
			for _, key := range keys {
				keyValue := viper.Get(key).(string)
				if keyValue == "" {
					keyValue = "[empty]"
				}
				fmt.Println(key + ": " + keyValue)
			}
		},
	}
)

// SetDefaultConfig returns the default configuration file structure
func SetDefaultConfig() error {

	user, err := user.Current()
	if err != nil {
		return err
	}

	home := user.HomeDir
	if err != nil {
		return errors.Wrap(err, "Failed to get home directory")
	}

	viper.Set(cConfigPath, home+string(os.PathSeparator)+".golinks"+string(os.PathSeparator)+"golinks.json")
	viper.Set(cTestPath, home+string(os.PathSeparator)+".golinks"+string(os.PathSeparator)+"test")
	viper.Set(cTempPath, home+string(os.PathSeparator)+".golinks"+string(os.PathSeparator)+"tmp")
	viper.Set(cStagingPath, home+string(os.PathSeparator)+".golinks"+string(os.PathSeparator)+"stage")

	if _, err := os.Stat(home + "/.golinks/golinks.json"); os.IsNotExist(err) {
		verb("creating golinks.json")
		if _, err := os.Create(home + "/.golinks/golinks.json"); err != nil {
			return err
		}
	}

	if err := viper.WriteConfig(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return nil
}
