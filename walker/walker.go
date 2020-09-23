package walker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//Walker contains the structure for a file walker
type Walker struct {
	workers int
	root    string
	archive []string
}

//New returns a new Walker
func New(root string) Walker {
	return Walker{1, root, nil}
}

//Workers returns the number of current workers
func (w Walker) Workers() int {
	return w.workers
}

//Root returns the current walker root
func (w Walker) Root() string {
	return w.root
}

//Archive returns the walkers archive if set
func (w Walker) Archive() []string {
	return w.archive
}

//PrintArchive prints all files in the existing archive
func (w Walker) PrintArchive() {
	if len(w.archive) == 0 {
		fmt.Println("archive empty")
		return
	}
	for _, r := range w.archive {
		fmt.Printf("%s\n", r)
	}
}

//Walk handles walking of a walkers root filesystem. Inaccessable directories are skipped.
func (w *Walker) Walk() error {
	if w.root == "" {
		return errors.New("Walk: Archive Empty")
	}
	e := filepath.Walk(w.root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		if strings.Contains(path, "Docker.raw") {
			return nil
		}
		if !f.IsDir() && f.Mode().IsRegular() {
			f, err := os.Open(path)
			if !os.IsPermission(err) {
				w.archive = append(w.archive, path)
			}
			f.Close()
		}
		return err
	})
	return e
}
