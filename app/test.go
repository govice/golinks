package main

import (
	"fmt"
	"math/rand"

	"io/ioutil"
	"time"

	"github.com/urfave/cli"
)

const (
	smallArchive int = 10 // Small File Archive
	smallDir     int = 10
	smallFile    int = 1000 //bytes

)

var testPath = ""

type test struct {
	ArchiveSize   int
	DirectorySize int
	FileSize      int
}

//TODO cleanup of testing environment
func appBuildTestDir(c *cli.Context) error {
	//fmt.Println("building test directory")
	fmt.Println(c.FlagNames())
	fmt.Println(testPath)

	if err := verifyPath(c.Args().First()); err != nil {
		return err
	}
	testPath = c.Args().First()

	for _, flagName := range c.FlagNames() {
		switch flagName {
		case "small":
			small := &test{smallArchive, smallDir, smallFile}
			fmt.Println("generating small test")
			if err := generateTestDir(testPath, small); err != nil {
				return cli.NewExitError("BuildTestDir: failed to build small test directory", 0)
			}

		default:
			fmt.Println("default called")
		}
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
