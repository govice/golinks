package blockmap

import (
	"path/filepath"

	"bytes"
	"crypto/sha512"

	"encoding/gob"

	"fmt"

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

	//If we're here, the entries are successful so we'll hash the blockmap.
	if err := b.hashBlockMap(); err != nil {
		return errors.Wrap(err, "BlockMap: failed to has block map")
	}

	return nil

}

func (b *BlockMap) hashBlockMap() error {
	if b.archive == nil {
		return errors.New("hashBlockMap: Attempted to hash null archive")
	}
	hash := sha512.New()
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(b.archive)
	if err != nil {
		return errors.Wrap(err, "hashBlockMap: failed to encode archive map")
	}

	if _, err := hash.Write(buffer.Bytes()); err != nil {
		return errors.Wrap(err, "hashBlockMap: failed to write to write hash buffer")
	}

	b.rootHash = hash.Sum(nil)
	return nil

}

//PrintBlockMap prints an existing block map and returns an error if not configured
func (b BlockMap) PrintBlockMap() {
	if b.rootHash == nil {
		fmt.Println("BlockMap is unhashed or unset")
	}
	fmt.Println("Root: " + b.root)
	fmt.Printf("Hash: %v\n", b.rootHash)
	for key, value := range b.archive {
		fmt.Printf("%v: %v\n", key, value)
	}
}
