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
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var buildTestCmd = &cobra.Command{
	Use:   "buildtest",
	Short: "Build out an archive to test on",
	Long:  "Build out an archive to test on based according to the preferred size",
	Run: func(cmd *cobra.Command, args []string) {
		if err := buildTestDir(cmd.Flag("size").Value.String()); err != nil {
			log.Println(err)
			cmd.Help()
		}
	},
}

var cleanTestCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean test directory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cleanTestDir(); err != nil {
			log.Println(err)
			cmd.Help()
		}

	},
}

const (
	testDirs    int = 10
	testDirSize int = 10
	smallFile   int = 1                 //1 B
	mediumFile  int = smallFile * 1024  //1 KB
	largeFile   int = mediumFile * 1024 //1 MB
)

var (
	testSize  string
	randomize bool
)

type test struct {
	ArchiveSize   int
	DirectorySize int
	FileSize      int
}

func buildTestDir(size string) error {
	var testConfig test
	switch size {

	case "small":
		testConfig = test{testDirs, testDirSize, smallFile}
		verb("generating small test")

	case "medium":
		testConfig = test{testDirs, testDirSize, mediumFile}
		verb("generating medium test")

	case "large":
		testConfig = test{testDirs, testDirSize, largeFile}
		verb("generating large test")
	default:
		return errors.New("Build test size is not recognized")
	}

	testPath := viper.Get(cTestPath).(string)
	os.Mkdir(testPath, 0755)
	if good, err := verifyPath(testPath); !good || err != nil {
		if err != nil {
			return err
		}
		return errors.New("Failed to verify config test directory at " + testPath)
	}

	if err := generateTestDir(testPath, testConfig); err != nil {
		return err
	}
	return nil
}

func generateTestDir(testRoot string, t test) error {
	verb("Test Archive Size: " + strconv.Itoa(t.ArchiveSize))
	verb("Test Root: " + testRoot)
	for i := 0; i < t.ArchiveSize; i++ {
		iStr := strconv.Itoa(i)

		verb("Creating Test Archive " + iStr)
		testDir := testRoot + string(os.PathSeparator) + "test" + iStr
		if err := os.Mkdir(testDir, 0644); err != nil {
			return err
		}

		//Write random data to files
		var source rand.Source
		if randomize {
			source = rand.NewSource(time.Now().UnixNano())
		} else {
			source = rand.NewSource(1)
		}
		r := rand.New(source)
		for j := 0; j < t.DirectorySize; j++ {
			jStr := strconv.Itoa(j)
			buff := make([]byte, t.FileSize)
			r.Read(buff)
			testFile := testDir + string(os.PathSeparator) + "file" + jStr
			if err := ioutil.WriteFile(testFile, buff, 0644); err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func verifyPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func cleanTestDir() error {
	testPath := viper.Get(cTestPath).(string)
	verb("removing " + testPath)
	if err := os.RemoveAll(testPath); err != nil {
		return err
	}
	return nil
}
