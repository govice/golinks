package main

import (
	"fmt"
	"math/rand"

	"io/ioutil"
	"time"

	"os"

	"github.com/urfave/cli"
)

const (
	testDirs    int = 10
	testDirSize int = 10
	smallFile   int = 1                 //1 B
	mediumFile  int = smallFile * 1024  //1 KB
	largeFile   int = mediumFile * 1024 //1 MB
)

var testPath = "/testing"

type test struct {
	ArchiveSize   int
	DirectorySize int
	FileSize      int
}

//TODO cleanup of testing environment
func appBuildTestDir(c *cli.Context) error {
	fmt.Println(c.FlagNames())
	fmt.Println(testPath)

	if err := verifyPath(c.Args().First()); err != nil {
		return err
	}

	testPath = c.Args().First() + testPath
	if err := os.Mkdir(testPath, 0644); err != nil {
		return cli.NewExitError("failed to create test directory", 0)
	}

	var config *test
	for _, flagName := range c.FlagNames() {
		if !c.IsSet(flagName) {
			continue
		}
		switch flagName {

		case "small":
			config = &test{testDirs, testDirSize, smallFile}
			fmt.Println("generating small test")

		case "medium":
			config = &test{testDirs, testDirSize, mediumFile}
			fmt.Println("generating medium test")

		case "large":
			config = &test{testDirs, testDirSize, largeFile}
			fmt.Println("generating large test")

		default:
			continue
		}
	}
	if err := generateTestDir(testPath, config); err != nil {
		return cli.NewExitError("BuildTestDir: failed to build small test directory", 0)
	}

	return nil
}

func generateTestDir(testRoot string, t *test) error {
	fmt.Println("generating test directory")
	for i := 0; i < t.ArchiveSize; i++ {
		tmpdir, err := ioutil.TempDir(testPath, "testDir")
		if err != nil {
			return err
		}

		//Write random data to files
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for j := 0; j < t.DirectorySize; j++ {
			buff := make([]byte, t.FileSize)
			r.Read(buff)
			tmpfile, err := ioutil.TempFile(tmpdir, "testFile")
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
