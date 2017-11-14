package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//Walker contains the structure for a file walker
type Walker struct {
	workers int
	root    string
}

//New returns a new Walker
func New(root string) Walker {
	return Walker{1, root}
}

//Workers returns the number of current workers
func (w Walker) Workers() int {
	return w.workers
}

//Root returns the current walker root
func (w Walker) Root() string {
	return w.root
}

//Walk handles walking of a walkers root filesystem
func (w Walker) Walk() error {
	e := filepath.Walk(w.root, func(path string, f os.FileInfo, err error) error {
		files, _ := ioutil.ReadDir(path)
		for _, r := range files {
			fmt.Printf("Visited %s\n", r.Name())

		}
		return err
	})
	return e
}
