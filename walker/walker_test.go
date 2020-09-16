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

package walker

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const (
	SmallArchive int = 10 // Small File Archive
	SmallDir     int = 10
	SmallFile    int = 1000 //bytes

)

var TestPath, _ = filepath.Abs(filepath.Dir(os.Args[0]) + "/testHome/") //Testing Root

func TestWalker_Walker(t *testing.T) {
	log.Println("Testing Walker")

	//setup

	//create files to test
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := os.Mkdir(TestPath, 0755); err != nil {
		t.Error(err)
	}
	//remove files after testing
	defer func() {
		if err := os.RemoveAll(TestPath); err != nil {
			t.Error(err)
		}
	}()
	//Generate Archive
	for i := 0; i < SmallArchive; i++ {
		tmpdir, err := ioutil.TempDir(TestPath, "testDir")
		if err != nil {
			t.Error(err)
		}

		//Write random data to files
		for j := 0; j < SmallDir; j++ {
			buff := make([]byte, SmallFile)
			r.Read(buff)
			tmpfile, err := ioutil.TempFile(tmpdir, "testFile")
			if err != nil {
				t.Error(err)
			}

			if _, err := tmpfile.Write(buff); err != nil {
				t.Error(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Error(err)
			}
		}
	} //end setup

	//Run Tests
	t.Run("Walker Walk", func(t *testing.T) {
		log.Println("Testing Walker Walk")
		w := New(TestPath)
		if err := w.Walk(); err != nil {
			t.Error(err)
		}
		//w.PrintArchive()
	})

	t.Run("Walk non-permissive", func(t *testing.T) {
		log.Println("testing walker non-permissive")

		nonperm := filepath.Join(TestPath, "nonpermissive")
		if err := os.Mkdir(nonperm, 0755); err != nil {
			t.Error(err)
		}

		//Write random data to files
		for j := 0; j < SmallDir; j++ {
			buff := make([]byte, SmallFile)
			r.Read(buff)
			tmpfile, err := ioutil.TempFile(nonperm, "testFile")
			if err != nil {
				t.Error(err)
			}

			if _, err := tmpfile.Write(buff); err != nil {
				t.Error(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Error(err)
			}
		}

		// make directory nonpermissive
		if err := os.Chmod(nonperm, 0000); err != nil {
			t.Error(err)
		}

		w := New(nonperm)
		if err := w.Walk(); err != nil {
			os.Chmod(nonperm, 0755)
			t.Error("expected skip for non-permissive directory")
		}
		os.Chmod(nonperm, 0755)

		if len(w.archive) != 0 {
			t.Error("archive length", len(w.archive), "does not match expected:", 0)
		}
	})
}
