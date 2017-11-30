package fs

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

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

func TestCompress(t *testing.T) {
	if err := Compress(os.Getenv("TEST_ARCHIVE")); err != nil {
		t.Error(err)
	}
}
