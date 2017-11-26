package blockmap

import (
	"log"

	"path/filepath"

	"github.com/LaughingCabbage/goLinks/types/fs"
	"github.com/LaughingCabbage/goLinks/types/walker"
	"github.com/pkg/errors"
)

//BlockMap is a ad-hoc Merkle tree-map
type BlockMap struct {
	archive  map[string][]byte
	rootHash []byte
	root     string
}

//New returns a new BlockMap initialized at the provided root
func New(root string) *BlockMap {
	rootMap := make(map[string][]byte)
	return &BlockMap{archive: rootMap, rootHash: nil, root: root}
}

//Generate creates an archive of the provided archives root filesystem
func (b *BlockMap) Generate() error {
	log.Println("Generating at root " + b.root)

	w := walker.New(b.root)
	if err := w.Walk(); err != nil {
		return errors.Wrap(err, "BlockMap: failed to walk "+w.Root())
	}

	//Iterate through all walked files
	for _, filePath := range w.Archive() {
		//Get the hash for the file
		fileHash, err := fs.HashFile(filePath)
		if err != nil {
			return errors.Wrap(err, "BlockMap: failed to hash "+filePath)
		}

		//Extract the relative path for the archive
		relPath, err := filepath.Rel(w.Root(), filePath)
		if err != nil {
			return errors.Wrap(err, "BlockMap: failed to extract relative file path")
		}

		//Add the hash to the archive using the relative path as it's key
		b.archive[relPath] = fileHash
	}
	return nil

}
