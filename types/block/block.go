//Package block provides an interface to create and maintain blockchain blocks.
package block

import (
	"bytes"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"log"
	"time"
)

//Block type structures a blockchain block by encapsulating vital components including index, timestamp, payload, and hashes.
type Block struct {
	Index      int
	Timestamp  time.Time
	Data       string
	Parenthash hash.Hash
	Blockhash  hash.Hash
}

//New creates a new blockchain block and initializes index, payload data, and hashes.
func New(index int, data string, parent hash.Hash) Block {
	blk := Block{index, time.Now(), data, nil, nil}
	//handle genesis block case
	if index == 0 {
		parenthash := sha512.New()
		parenthash.Write([]byte(blk.Data))
		//blk.Parenthash = parenthash.Sum(nil)
	} else {
		blk.Parenthash = parent
	}
	blk.Blockhash = blk.Hash()
	return blk
}

//Hash calculates and returns a SHA512 hash for the block.
func (block Block) Hash() hash.Hash {
	blkhash := sha512.New()
	buff, err := block.MarshalBinary()
	if err != nil {
		log.Fatal("Failed to hash block")
	}
	blkhash.Write(buff)
	return blkhash
}

//Validate compares two blocks to verify their parent child relationship.
func Validate(prev, current Block) error {
	if prev.Index+1 != current.Index {
		return errors.New("block indexes do not correlate")
	}
	if !bytes.Equal(prev.Blockhash.Sum(nil), current.Parenthash.Sum(nil)) {
		return errors.New("block parent child hashes do not correlate")
	}
	h := current.Hash().Sum(nil)
	if !bytes.Equal(h, current.Blockhash.Sum(nil)) {
		return errors.New("current block's hash is not valid")
	}
	return nil
}

//MarshalBinary serializes a block structure and returns the blocks byte sequence.
func (block Block) MarshalBinary() ([]byte, error) {
	var buffer bytes.Buffer
	fmt.Fprintln(&buffer, block.Index, block.Timestamp, block.Data, block.Parenthash)
	return buffer.Bytes(), nil
}

//UnmarshalBinary decodes a byte buffer into a block.
func (block *Block) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(buffer, &block.Index, &block.Timestamp, &block.Data, block.Parenthash)
	return err
}
