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

package blockchain

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/govice/golinks/block"

	"fmt"

	"github.com/pkg/errors"
)

//Blockchain type implements an array of blocks.
// type Blockchain []block.Block
type Blockchain struct {
	Blocks []block.Block `json:"blocks"`
}

type Blockchainer interface {
	New() Blockchain
	Add(b block.Block)
	Validate() error
	Length()
}

const genesisSize int = 100 //bytes

var ErrInvalidGenesisBlock = errors.New("blockchain: invalid genesis block")

//New returns a new blockchain and initializes the chain's genesis block.
func New(genesisBlock *block.Block) (*Blockchain, error) {
	blkchain := &Blockchain{}
	if len(genesisBlock.ParentHash) != 0 {
		return blkchain, ErrInvalidGenesisBlock
	}

	blkchain.Blocks = append(blkchain.Blocks, *genesisBlock)
	return blkchain, nil
}

func (b *Blockchain) Length() int {
	return len(b.Blocks)
}

func (b *Blockchain) At(index int) *block.Block {
	return &b.Blocks[index]
}

//AddSHA512 adds a new block to the chain given a payload.
func (b *Blockchain) AddSHA512(data []byte) *block.Block {
	blk := block.NewSHA512(b.Length(), data, b.Blocks[b.Length()-1].BlockHash)
	b.Blocks = append(b.Blocks, *blk)
	return blk
}

//Print outputs the blockchain to standard output.
func (b *Blockchain) Print() {
	for i := 0; i < len(b.Blocks); i++ {
		fmt.Println("Block ", i, ": ", b.Blocks[i])
	}
}

//Validate iterates through blocks and calls the block.validate method for the length of the chain.
func (b *Blockchain) Validate() error {
	if b.Length() < 2 {
		return errors.New("Validate: invalid genesis block")
	}
	for i := 1; i < b.Length(); i++ {
		if err := block.Validate(b.At(i-1), b.At(i)); err != nil {
			return errors.Wrap(err, "Validate: failed to validate blockchain blocks")
		}
	}
	return nil
}

//GetCurrentHash Returns the most recent hash in a blockchain
func (b *Blockchain) GetCurrentHash() []byte {
	return b.Blocks[b.Length()].BlockHash
}

// SubChain returns a new blockchain at index of blockchain
func (b *Blockchain) SubChain(index int) (*Blockchain, error) {
	if index == b.Length() {
		return b, nil
	}

	chain := &Blockchain{}

	if index > b.Length() {
		return chain, errors.New("blockchain: Subchain index exceeds chain length")
	}

	chain.Blocks = append(chain.Blocks, b.Blocks[:index]...)
	return chain, nil
}

//GetGCI returns the greatest common index between the current blockchain and the new blockchain
func (b *Blockchain) GetGCI(other *Blockchain) (int, error) {
	gci := b.Length()
	bSubchain := &Blockchain{}
	otherSubchain := &Blockchain{}
	var err error
	if other.Length() > b.Length() {
		otherSubchain, err = other.SubChain(b.Length())
		if err != nil {
			return -1, err
		}
		bSubchain = b
	} else if b.Length() > other.Length() {
		gci = other.Length()
		bSubchain, err = b.SubChain(other.Length())
		if err != nil {
			return -1, err
		}
		otherSubchain = other
	} else {
		bSubchain = b
		otherSubchain = other
	}

	if !Equal(bSubchain, otherSubchain) {
		return -1, errors.New("blockchain: blockchains are not equal")
	}

	return gci, nil
}

//Equal tests the equality of two blockchains
func Equal(chainA, chainB *Blockchain) bool {
	if chainA.Length() != chainB.Length() {
		return false
	}

	for i := 0; i < chainA.Length(); i++ {
		if !bytes.Equal(chainA.At(i).BlockHash, chainB.At(i).BlockHash) {
			return false
		}
	}
	return true
}

//Save saves the blockchain to a .dat file
func (blockchain Blockchain) Save(name string) error {
	file, err := os.OpenFile(name+".dat", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return errors.Wrap(err, "Save: failed to open file")
	}
	encoder := json.NewEncoder(file)
	if err = encoder.Encode(blockchain); err != nil {
		return errors.Wrap(err, "Save: failed to encode blockchain")
	}
	if err = file.Close(); err != nil {
		return errors.Wrap(err, "Save: failed to close file")
	}
	return err
}

//Load loads a blockchain from a .dat file and initializes the blockchain
func (blockchain *Blockchain) Load(name string) error {
	file, err := os.Open(name + ".dat")
	if err != nil {
		return errors.Wrap(err, "Load: failed to open file")
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(blockchain); err != nil {
		return errors.Wrap(err, "Load: failed to decode blockchain")
	}
	if err = file.Close(); err != nil {
		return errors.Wrap(err, "Load: failed to close file")
	}
	return err
}

func (b *Blockchain) FindByBlockHash(hash []byte) *block.Block {
	for _, block := range b.Blocks {
		if bytes.Equal(block.BlockHash, hash) {
			return &block
		}
	}

	return nil
}

func (b *Blockchain) FindByParentHash(hash []byte) *block.Block {
	for _, block := range b.Blocks {
		if bytes.Equal(block.ParentHash, hash) {
			return &block
		}
	}
	return nil
}

func (b *Blockchain) FindByTimestamp(timestamp int64) *block.Block {
	for _, block := range b.Blocks {
		if block.Timestamp == timestamp {
			return &block
		}
	}
	return nil
}

func Copy(other *Blockchain) *Blockchain {
	newChain := &Blockchain{
		Blocks: make([]block.Block, len(other.Blocks)),
	}
	copy(newChain.Blocks, other.Blocks)
	return newChain
}

var ErrNoCommonGCI = errors.New("blocks do not share a GCI")
var ErrShorterChain = errors.New("new chain is shorter than original")

// UpdateChain compares original and new chain and returns a copy of new
// iff orginal and new share a GCI and new is longer than original.
func UpdateChain(original, new *Blockchain) (*Blockchain, error) {
	//Chain is longer and needs updating.
	if _, err := original.GetGCI(new); err != nil {
		return nil, err
	}
	if new.Length() < original.Length() {
		if err := new.Validate(); err != nil {
			return nil, err
		}

		return nil, ErrShorterChain
	}

	chainCopy := *new
	return &chainCopy, nil
}
