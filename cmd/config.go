package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "handle tasks related to configuration",
		Run: func(cmd *cobra.Command, args []string) {
			configPath := viper.Get("configpath").(string)
			fmt.Println("config path: " + configPath)
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

//Config contains the structure used to create a configuration file
type Config struct {
	Name       string `json:"name"`
	TestPath   string `json:"testpath"`
	ConfigPath string `json:"configpath"`
}

// DefaultConfig returns the default configuration file structure
func DefaultConfig() (Config, error) {
	var config Config
	home, err := homedir.Dir()
	if err != nil {
		return config, errors.Wrap(err, "Failed to get home directory")
	}
	config = Config{
		ConfigPath: home + string(os.PathSeparator) + ".golinks" + string(os.PathSeparator) + "golinks.json",
		TestPath:   home + string(os.PathSeparator) + ".golinks" + string(os.PathSeparator) + "test",
	}
	return config, nil
}

// WriteConfig writes a config file to the path defined in Config.ConfigPath
func (c Config) WriteConfig() error {
	configJSON, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal config json")
	}

	if err := ioutil.WriteFile(c.ConfigPath, configJSON, 0644); err != nil {
		return errors.Wrap(err, "Failed to write config file")
	}
	return nil
}

// ReadConfig TODO see if needed
func ReadConfig(path string) (Config, error) {
	var config = Config{}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return config, errors.Wrap(err, "Failed to read config file")
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return config, errors.Wrap(err, "Failed to unmarshal config json")
	}

	return config, nil
}
