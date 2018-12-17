/*
 *Copyright 2018 Kevin Gentile
 *
 *Licensed under the Apache License, Version 2.0 (the "License");
 *you may not use this file except in compliance with the License.
 *You may obtain a copy of the License at
 *
 *http://www.apache.org/licenses/LICENSE-2.0
 *
 *Unless required by applicable law or agreed to in writing, software
 *distributed under the License is distributed on an "AS IS" BASIS,
 *WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *See the License for the specific language governing permissions and
 *limitations under the License.
 */

package block

import (
	"github.com/pkg/errors"

	"bytes"
	"crypto/sha512"
	"fmt"
	"time"
)

//Block type structures a blockchain block by encapsulating vital components including index, timestamp, payload, and hashes.
type Block struct {
	Index      int    `json:"index"`
	Timestamp  int64  `json:"timestamp"`
	Data       []byte `json:"data"`
	Parenthash []byte `json:"parentHash"`
	Blockhash  []byte `json:"blockhash"`
}

//New creates a new blockchain block and initializes index, payload data, and hashes.
func New(index int, data []byte, parent []byte) Block {
	blk := Block{index, time.Now().Unix(), data, nil, nil}
	//handle genesis block case
	if index == 0 {
		parenthash := sha512.New()
		if _, err := parenthash.Write(blk.Data); err != nil {
			panic("Failed to write hash")
		}
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
	if _, err := blkhash.Write(buffer.Bytes()); err != nil {
		panic("failed to write to hash buffer")
	}
	return blkhash.Sum(nil)
}

//Validate compares two blocks to verify their parent child relationship.
func Validate(prev, current Block) error {
	if prev.Index+1 != current.Index {
		return errors.New("Validate: block indexes do not correlate")
	}
	if !bytes.Equal(prev.Blockhash, current.Parenthash) {
		return errors.New("Validate: block hashes do not match")
	}
	h := current.Hash()
	if !bytes.Equal(h, current.Blockhash) {
		return errors.New("Validate: current block's hash is not valid")
	}
	return nil
}

//PrintBlock prints a block's members to a new line.
func (block Block) PrintBlock() {
	fmt.Println("Block: ", block.Index, block.Timestamp, block.Data, block.Parenthash, block.Blockhash)
}

//Equal checks if the hash of two blocks are equal.
func Equal(blockA, blockB Block) bool {
	blockAHash := blockA.Hash()
	blockBHash := blockB.Hash()
	return bytes.Equal(blockAHash, blockBHash)
}
