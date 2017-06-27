//Package block provides an interface to create and maintain blockchain blocks.
package block

import (
	"bytes"
	"crypto/sha512"
	"errors"
	"fmt"
	"time"
)

//Block type structures a blockchain block by encapsulating vital components including index, timestamp, payload, and hashes.
type Block struct {
	Index      int
	Timestamp  time.Time
	Data       string
	Parenthash []byte
	Blockhash  []byte
}

//New creates a new blockchain block and initializes index, payload data, and hashes.
func New(index int, data string, parent []byte) Block {
	blk := Block{index, time.Now(), data, nil, nil}
	//handle genesis block case
	if index == 0 {
		parenthash := sha512.New()
		parenthash.Write([]byte(blk.Data))
		blk.Parenthash = parenthash.Sum(nil)
	} else {
		blk.Parenthash = parent
	}
	blk.Blockhash = blk.Hash()
	return blk
}

//Hash calculates and returns a SHA512 hash for the block.
func (block Block) Hash() []byte {
	blkhash := sha512.New()
	var buffer bytes.Buffer
	fmt.Fprintln(&buffer, block.Index, block.Timestamp, block.Data, block.Parenthash)
	blkhash.Write(buffer.Bytes())
	return blkhash.Sum(nil)
}

//Validate compares two blocks to verify their parent child relationship.
func Validate(prev, current Block) error {
	if prev.Index+1 != current.Index {
		return errors.New("block indexes do not correlate")
	}
	if !bytes.Equal(prev.Blockhash, current.Parenthash) {
		return errors.New("block parent child hashes do not correlate")
	}
	h := current.Hash()
	if !bytes.Equal(h, current.Blockhash) {
		return errors.New("current block's hash is not valid")
	}
	return nil
}

//PrintBlock prints a block's members to a new line.
func (block Block) PrintBlock() {
	fmt.Println("Block: ", block.Index, block.Timestamp, block.Data, block.Parenthash, block.Blockhash)
}

//Equal checks if the hash of two blocks are equal.
func Equal(blockA, blockB Block) bool {
	if !bytes.Equal(blockA.Hash(), blockB.Hash()) {
		return false
	}
	return true
}
