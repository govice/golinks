package fs

import (
	"fmt"
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
	err := filepath.Walk(w.root, visit)
	return err
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited %s\n", path)
	return nil
}
