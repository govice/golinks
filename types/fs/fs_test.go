/*
 *Copyright 2017 Kevin Gentile
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

package fs

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/laughingcabbage/golinks/types/blockchain"
	"github.com/pkg/errors"
)

func TestHashFile(t *testing.T) {
	//create files to test
	var TestPath, _ = filepath.Abs(filepath.Dir(os.Args[0]) + "/testHome/") //Testing Root
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := os.Mkdir(TestPath, 0644); err != nil {
		t.Error(err)
	}
	//remove files after testing

	defer func() {
		if err := os.RemoveAll(TestPath); err != nil {
			t.Error(err)
		}
	}()

	//Generate Archive
	tmpdir, err := ioutil.TempDir(TestPath, "testDir")
	if err != nil {
		t.Error(err)
	}

	//Write random data to files
	buff := make([]byte, 1000000)
	r.Read(buff)
	tmpfile, err := ioutil.TempFile(tmpdir, "testFile")
	if err != nil {
		t.Error(err)
	}

	if _, err = tmpfile.Write(buff); err != nil {
		t.Error(err)
	}

	if err = tmpfile.Close(); err != nil {
		t.Error(err)
	}
	//end setup

	//Run hash
	hash, err := HashFile(tmpfile.Name())
	if err != nil {
		t.Error(err)
	}
	if hash == nil {
		t.Error(errors.New("fs: HashFile returned nil hash"))
	}

}

func TestZip(t *testing.T) {
	t.SkipNow()
	if err := Compress(os.Getenv("TEST_FOLDER"), os.Getenv("ZIP_DEST")); err != nil {
		t.Error(err)
	}
	if err := Decompress(os.Getenv("TEST_ARCHIVE"), os.Getenv("ZIP_DEST")); err != nil {
		t.Error(err)
	}

	//TODO move this test to blockmap (import cycle)
	/*
			original := &blockmap.BlockMap{}
			if err := original.Load(os.Getenv("TEST_FOLDER")); err != nil {
				t.Error(err)
			}

			unziped := &blockmap.BlockMap{}
			if err := unziped.Load(os.Getenv("ZIP_OUT")); err != nil {
				t.Error(err)
			}

		//Test validity
		if !blockmap.Equal(original, unziped) {
			t.Error("Original and unziped archives are different")
		}
	*/

	//Cleanup
	if err := os.RemoveAll(os.Getenv("ZIP_OUT")); err != nil {
		t.Error(err)
	}

}

//todo create test folder
func TestSaveGob(t *testing.T) {
	fileName := "testFile.link"
	cwd, _ := os.Getwd()

	filePath := cwd + string(os.PathSeparator) + fileName
	fmt.Println(filePath)
	b := blockchain.New()
	if err := SaveGob(filePath, b); err != nil {
		t.Error(err)
	}
	defer os.Remove(filePath)

	blockChainOut := blockchain.Blockchain{}
	ReadGob(filePath, &blockChainOut)

	if !blockchain.Equal(b, blockChainOut) {
		t.Error("Gob file not written/read properly")
	}
}
