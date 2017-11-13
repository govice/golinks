package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

const (
	SmallDir  = 10
	SmallFile = 1000 //bytes
)

//TODO finish test
func TestWalker_Walker(t *testing.T) {
	//setup
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(dir)
	}

	for i := 0; i < SmallDir; i++ {
		buff := make([]byte, SmallFile)
		index := strconv.Itoa(i)
		filename := "test" + index + ".link"
		if err := ioutil.WriteFile(dir+filename, buff, 0666); err != nil {
			t.Error(err)
		}
	}

	//Run Tests
	t.Run("TestWalk", TestWalker_Walk)

}

func TestWalker_Walk(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Error(err)
	}

	w := New(dir + "/test")
	if err := w.Walk(); err != nil {
		t.Error(err)
	}
}
