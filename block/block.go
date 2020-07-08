/*
 *Copyright 2018-2019 Kevin Gentile
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
	"bytes"
	"crypto/sha512"
	"encoding/json"
	"hash"
	"time"

	"github.com/pkg/errors"
)

// Blocker the interface used to implement a block
type Blocker interface {
	// Serialize marshals a block into JSON for hashing, omitting the hash BlockHash field
	Serialize() ([]byte, error)

	// Hash block after seralization
	Hash(hasher hash.Hash) ([]byte, error)
}

// Block describes a block for use in a blockchain
type Block struct {
	Index      int    `json:"index"`
	Timestamp  int64  `json:"timestamp"`
	Data       []byte `json:"data"`
	ParentHash []byte `json:"parentHash"`
	BlockHash  []byte `json:"blockHash,omitempty"`
}

// NewSHA512 creates a new block using SHA512 hashing and generates its hash
func NewSHA512(index int, data []byte, parentHash []byte) *Block {
	blk := &Block{
		Index:      index,
		Timestamp:  time.Now().UnixNano(),
		Data:       append([]byte{}, data...),
		ParentHash: append([]byte{}, parentHash...),
	}
	blk.Hash(sha512.New())
	return blk
}

// NewSHA512Genesis returns a new gensis block hashed with SHA512
func NewSHA512Genesis() *Block {
	genesis := &Block{
		Index:      0,
		Timestamp:  time.Time{}.UnixNano(),
		Data:       append([]byte{}, []byte("GENSIS BLOCK")...),
		ParentHash: append([]byte{}, []byte("")...),
	}
	genesis.Hash(sha512.New())
	return genesis
}

// Hash returns the block's hash
func (block *Block) Hash(hasher hash.Hash) ([]byte, error) {
	// Hash if not already hashed
	if len(block.BlockHash) == 0 {
		blockBytes, err := block.Serialize()
		if err != nil {
			return nil, err
		}
		if _, err := hasher.Write(blockBytes); err != nil {
			return nil, err
		}
		block.BlockHash = append([]byte{}, hasher.Sum(nil)...)
	}
	return block.BlockHash, nil
}

// Serialize returns json bytes for hashing
func (block *Block) Serialize() ([]byte, error) {
	jsonBlock := &Block{
		Index:      block.Index,
		Timestamp:  block.Timestamp,
		ParentHash: append([]byte{}, block.ParentHash...),
		Data:       append([]byte{}, block.Data...),
	}
	jsonBytes, err := json.Marshal(jsonBlock)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// ErrBadParentChild is returned for an invalid block validation
var ErrBadParentChild = errors.New("block: invalid parent-child relationship")

//Validate compares two blocks to verify their parent child hash relationship.
func Validate(prev, current *Block) error {
	if prev.Index+1 != current.Index {
		return ErrBadParentChild
	}
	if !bytes.Equal(prev.BlockHash, current.ParentHash) {
		return ErrBadParentChild
	}

	if !bytes.Equal(prev.BlockHash, current.ParentHash) {
		return ErrBadParentChild
	}
	return nil
}

// Equal returns true if block is equal to other
func Equal(block, other *Block) bool {
	if block.Index != other.Index {
		return false
	}

	if !bytes.Equal(block.BlockHash, other.BlockHash) {
		return false
	}

	if !bytes.Equal(block.ParentHash, other.ParentHash) {
		return false
	}

	if block.Timestamp != other.Timestamp {
		return false
	}

	if !bytes.Equal(block.Data, other.Data) {
		return false
	}

	return true
}
