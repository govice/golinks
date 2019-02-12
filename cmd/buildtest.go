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

var testSize string

type test struct {
	ArchiveSize   int
	DirectorySize int
	FileSize      int
}

//TODO cleanup of testing environment
func buildTestDir(size string) error {
	var testConfig test
	switch size {

	case "small":
		testConfig = test{testDirs, testDirSize, smallFile}
		log.Print("generating small test")

	case "medium":
		testConfig = test{testDirs, testDirSize, mediumFile}
		log.Println("generating medium test")

	case "large":
		testConfig = test{testDirs, testDirSize, largeFile}
		log.Println("generating large test")
	default:
		return errors.New("Build test size is not recognized")
	}

	testPath := viper.Get("testpath").(string)
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

// TODO maybe not random data?
func generateTestDir(testRoot string, t test) error {
	log.Printf("Test Archive Size: %v", t.ArchiveSize)
	log.Println("Test Root: " + testRoot)
	for i := 0; i < t.ArchiveSize; i++ {
		iStr := strconv.Itoa(i)
		log.Println("Creating Archive " + iStr)
		tmpdir, err := ioutil.TempDir(testRoot, "test"+iStr)
		if err != nil {
			return err
		}
		//Write random data to files
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for j := 0; j < t.DirectorySize; j++ {
			jStr := strconv.Itoa(j)
			buff := make([]byte, t.FileSize)
			r.Read(buff)
			tmpfile, err := ioutil.TempFile(tmpdir, "file"+jStr)
			if err != nil {
				panic(err)
			}

			if _, err := tmpfile.Write(buff); err != nil {
				panic(err)
			}
			if err := tmpfile.Close(); err != nil {
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
	testPath := viper.Get("testpath").(string)
	log.Println("removing " + testPath)
	if err := os.RemoveAll(testPath); err != nil {
		return err
	}
	return nil
}
