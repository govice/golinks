//Package blockchain provides handles to create and maintain a blockchain
package blockchain

import (
	"github.com/LaughingCabbage/goLinks/types/block"

	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
)

//Blockchain type implements an array of blocks.
type Blockchain []block.Block

//New returns a new blockchain and initializes the chain's genesis block.
func New() Blockchain {
	var blkchain Blockchain
	//create genesis block and append it as root to blockchain
	blk := block.New(0, "GENESIS DATA", nil)
	blkchain = append(blkchain, blk)
	return blkchain
}

//Add adds a new block to the chain given a payload.
func (blockchain *Blockchain) Add(data string) {
	blk := block.New(len(*blockchain), data, (*blockchain)[len(*blockchain)-1].Blockhash)
	*blockchain = append(*blockchain, blk)
}

//Print outputs the blockchain to standard output.
func (blockchain Blockchain) Print() {
	fmt.Println("Printing blockchain...")
	for i := 0; i < len(blockchain); i++ {
		fmt.Println("Block ", i, ": ", blockchain[i])
	}
}

//Validate iterates through blocks and calls the block.validate method for the length of the chain.
func (blockchain Blockchain) Validate() error {
	if len(blockchain) < 2 {
		return errors.New("invalid attempt to validate genesis block")
	}
	for i := 1; i < len(blockchain); i++ {
		err := block.Validate(blockchain[i-1], blockchain[i])
		if err != nil {
			return errors.New("blockchain is invalid")
		}
	}
	return nil
}

func (blockchain Blockchain) encodeChain() []byte {
	buffer := bytes.Buffer{}
	chainGob := gob.NewEncoder(&buffer)
	err := chainGob.Encode(blockchain)
	if err != nil {
		log.Fatal("failed to encode blockchain")
	}
	return buffer.Bytes()
}

func (blockchain *Blockchain) decodeChain(data []byte) error {
	buffer := bytes.Buffer{}
	buffer.Write(data)
	dec := gob.NewDecoder(&buffer)
	err := dec.Decode(blockchain)
	return err
}

//GetValidChain returns the longest valid chain given two blockchains.
// it should be implied that the longest chain should be the most recent valid chain
//this function should only take accept validated blockchains
func GetValidChain(current, new Blockchain) Blockchain {
	if len(new) > len(current) {
		return new
	}
	return current
}
