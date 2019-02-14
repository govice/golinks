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
type Blockchain []block.Block

type BlockchainJSON []block.BlockJSON

const genesisSize int = 100 //bytes

//New returns a new blockchain and initializes the chain's genesis block.
func New(genesisBlock block.Block) Blockchain {
	var blkchain Blockchain
	blkchain = append(blkchain, genesisBlock)
	return blkchain
}

//AddSHA512 adds a new block to the chain given a payload.
func (blockchain *Blockchain) AddSHA512(data []byte) {
	blk := block.NewSHA512(len(*blockchain), data, (*blockchain)[len(*blockchain)-1].Blockhash())
	*blockchain = append(*blockchain, blk)
}

//Print outputs the blockchain to standard output.
func (blockchain Blockchain) Print() {
	for i := 0; i < len(blockchain); i++ {
		fmt.Println("Block ", i, ": ", blockchain[i])
	}
}

//Validate iterates through blocks and calls the block.validate method for the length of the chain.
func (blockchain Blockchain) Validate() error {
	if len(blockchain) < 2 {
		return errors.New("Validate: invalid genesis block")
	}
	for i := 1; i < len(blockchain); i++ {
		if err := block.Validate(blockchain[i-1], blockchain[i]); err != nil {
			return errors.Wrap(err, "Validate: failed to validate blockchain blocks")
		}
	}
	return nil
}

//GetCurrentHash Returns the most recent hash in a blockchain
func (blockchain Blockchain) GetCurrentHash() []byte {
	return blockchain[len(blockchain)].Blockhash()
}

//UpdateChain returns the longest valid chain given two blockchains.
// it should be implied that the longest chain should have the most recent block
func (blockchain *Blockchain) UpdateChain(new Blockchain) error {
	//Chain is longer and needs updating.
	if blockchain.GetGCI(new) == -1 {
		return errors.New("UpdateChain: invalid GCI comparison")
	}
	if len(new) > len(*blockchain) {
		if err := new.Validate(); err != nil {
			return errors.Wrap(err, "UpdateChain: failed to validate new")
		}
		*blockchain = new
		return nil
	}
	return errors.New("UpdateChain: Failed")
}

//GetGCI returns the greatest common index between the current blockchain and the new blockchain
func (blockchain Blockchain) GetGCI(new Blockchain) int {
	if len(new) > len(blockchain) {
		if !Equal(blockchain, new[:len(blockchain)]) {
			return -1
		}
	}
	return len(blockchain)
}

//Equal tests the equality of two blockchains
func Equal(chainA, chainB Blockchain) bool {
	if len(chainA) != len(chainB) {
		return false
	}

	for i := 0; i < len(chainA); i++ {
		if !bytes.Equal(chainA[i].Blockhash(), chainB[i].Blockhash()) {
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

func (blockchain *Blockchain) MarshalJSON() ([]byte, error) {
	var blockchainJSON BlockchainJSON
	for _, blk := range *blockchain {
		blockchainJSON = append(blockchainJSON, blk.JSON())
	}

	chainBytes, err := json.Marshal(blockchainJSON)
	if err != nil {
		return nil, err
	}

	return chainBytes, nil
}

func (blockchain *Blockchain) UnmarshalJSON(bytes []byte) error {
	blockchainJSON := BlockchainJSON{}
	if err := json.Unmarshal(bytes, &blockchainJSON); err != nil {
		return err
	}

	for _, blockJSON := range blockchainJSON {
		*blockchain = append(*blockchain, blockJSON.Block())
	}
	return nil
}
