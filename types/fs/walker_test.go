package fs

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
	SmallDir  = 10
	SmallFile = 1000 //bytes
)

var TestPath, _ = filepath.Abs(filepath.Dir(os.Args[0]) + "/test/") //Testing Root

//TODO finish test
func TestWalker_Walker(t *testing.T) {
	log.Println("Testing Walker")

	//setup

	//create files to test
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
	//Write random data to files
	for i := 0; i < SmallDir; i++ {
		buff := make([]byte, SmallFile)
		r.Read(buff)
		//index := strconv.Itoa(i)
		tmpfile, err := ioutil.TempFile(TestPath, "test")
		if err != nil {
			t.Error(err)
		}

		if _, err := tmpfile.Write(buff); err != nil {
			t.Error(err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Error(err)
		}
	} // end setup

	//Run Tests
	t.Run("Walker Walk", func(t *testing.T) {
		log.Println("Testing Walker Walk")
		w := New(TestPath)
		if err := w.Walk(); err != nil {
			t.Error(err)
		}
	})

}
