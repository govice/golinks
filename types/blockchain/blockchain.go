//TODO create block tree

//Package blockchain provides handles to create and maintain a blockchain
package blockchain

import (
	"log"

	"github.com/LaughingCabbage/goLinks/types/block"

	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

//Blockchain type implements an array of blocks.
type Blockchain []block.Block

//New returns a new blockchain and initializes the chain's genesis block.
func New() Blockchain {
	var blkchain Blockchain
	//create genesis block and append it as root to blockchain
	blk := block.New(0, []byte("GENESIS DATA"), nil)
	blkchain = append(blkchain, blk)
	return blkchain
}

//Add adds a new block to the chain given a payload.
func (blockchain *Blockchain) Add(data []byte) {
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
		if err := block.Validate(blockchain[i-1], blockchain[i]); err != nil {
			return errors.New("blockchain is invalid")
		}
	}
	return nil
}

//GetCurrentHash Returns the most recent hash in a blockchain
func (blockchain Blockchain) GetCurrentHash() []byte {
	return blockchain[len(blockchain)].Blockhash
}

//UpdateChain returns the longest valid chain given two blockchains.
// it should be implied that the longest chain should be the most recent valid chain
//this function should only take accept validated blockchains
func (blockchain *Blockchain) UpdateChain(new Blockchain) error {
	//Chain is longer and needs updating.
	if blockchain.GetGCI(new) == 0 {
		return errors.New("invalid GCI comparison in UpdateChain")
	}
	if len(new) > len(*blockchain) {
		if err := new.Validate(); err != nil {
			return errors.New("failed to Update Chain. Chain Invalid")
		}
		*blockchain = new
		return nil
	}
	return errors.New("Chain not updated")
}

//GetGCI returns the greatest common index between the current blockchain and the new blockchain
func (blockchain Blockchain) GetGCI(new Blockchain) int {
	if len(new) > len(blockchain) {
		if !Equal(blockchain, new[:len(blockchain)]) {
			return 0
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
		if !block.Equal(chainA[i], chainB[i]) {
			log.Println("Chains not equal at blocks:")
			log.Println(chainA[i])
			log.Println(chainB[i])
			return false
		}
	}
	return true
}

//Save saves the blockchain to a .dat file
func (blockchain Blockchain) Save(name string) error {
	file, err := os.OpenFile(name+".dat", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(blockchain); err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
	//err := ioutil.WriteFile(filename, blockchain.encodeChain(), 0600)
	return err
}

//Load loads a blockchain from a .dat file and initializes the blockchain
func (blockchain *Blockchain) Load(name string) error {
	file, err := os.Open(name + ".dat")
	if err != nil {
		panic(err)
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(blockchain)
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
	return err
}
