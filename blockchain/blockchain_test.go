package blockchain

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/govice/golinks/block"
)

var genesisBlock = block.NewSHA512Genesis()

//TestGetValidChain creates two different blockchains of different sizes and attempts to validate the chain.
func TestBlockchain_Validate(t *testing.T) {
	blkchain, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	blkchain.AddSHA512([]byte("NewSTring"))
	blkchain.AddSHA512([]byte("NewSTring"))

	if err := blkchain.Validate(); err != nil {
		blkchain.Print()
		t.Error(err)
	}
	blkchain2, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	blkchain2.AddSHA512([]byte("chain2str"))
	blkchain2.AddSHA512([]byte("chain2str"))
	blkchain2.AddSHA512([]byte("data"))
	err = blkchain2.Validate()
	if err != nil {
		t.Error("Failed to Valiate Blockchain")
	}

	if Equal(blkchain, blkchain2) {
		t.Error("Validity check for blockchain binaries failed")
	}
}

func TestBlockchain_Equal(t *testing.T) {
	log.Println("Testing Equal")
	//construct two chains with genesis blocks
	chainA, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}

	//Test with equal blocks
	chainA.AddSHA512([]byte("NewSTring"))
	chainA.AddSHA512([]byte("NewSTring"))

	chainB := Copy(chainA)

	//Test equality with two equal blockchains
	if !Equal(chainA, chainB) {
		t.Error("equal chains are not equal")
	}

	//Test equality with additional block
	chainB.AddSHA512([]byte("data"))
	if Equal(chainA, chainB) {
		chainA.Print()

		chainB.Print()
		t.Error("unequal chains are testing as equal")

	}

}

func TestBlockchain_InputOutput(t *testing.T) {
	log.Println("Testing I/O")
	//Test saving to file
	blkchain, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	blkchain.AddSHA512([]byte("NewSTring"))
	blkchain.AddSHA512([]byte("NewSTring2"))

	if err := blkchain.Save("testfile"); err != nil {
		t.Error("failed to save blockchain ", err)
	}

	//Test loading from file
	blkchainB := &Blockchain{}
	err = blkchainB.Load("testfile")
	if err != nil {
		t.Error("failed to load blockchain", err)
	}

	//Test validity of read chain
	if !Equal(blkchain, blkchainB) {
		t.Error("read blockchain does not match saved chain")
	}

	//Cleanup test file
	err = os.Remove("testfile.dat")
	if err != nil {
		t.Error("failed to cleanup IO test file", err)
	}
}

func TestBlockchain_SubChain(t *testing.T) {
	log.Println("Testing SubChain")
	b, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))
	// Subchain of maximum length should return the same chain
	sub, err := b.SubChain(b.Length())
	if err != nil {
		t.Error(err)
	}

	if err != nil || !Equal(sub, b) {
		t.Errorf("Subchain %#v not equal to original chain: %#v", sub, b)
	}

	c := b
	c.AddSHA512([]byte("NewString3"))

	sub, err = c.SubChain(b.Length())
	if err != nil {
		t.Error(err)
	}
	if !Equal(sub, b) {
		t.Errorf("Subchain %#v not equal to original chain: %#v", sub, b)
	}
}

func TestBlockchain_GetGCI(t *testing.T) {
	log.Println("Testing GetGCI")
	b, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))
	c := b
	c.AddSHA512([]byte("new"))
	gci, err := b.GetGCI(c)
	if err != nil {
		t.Error(err)
	}
	if gci != b.Length() {
		t.Errorf("Invalid GCI of %v should be 3", gci)
	}

	gci, err = c.GetGCI(b)
	if err != nil {
		t.Error(err)
	}

	if gci != b.Length() {
		t.Errorf("Invalid GCI of %v should be 3", gci)
	}
}

func TestBlockchain_FindByHash(t *testing.T) {
	b, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	target := b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))

	result := b.FindByBlockHash(target.BlockHash)
	if result == nil {
		t.Fail()
	}

	if !block.Equal(target, result) {
		t.Fail()
	}

	result = b.FindByBlockHash([]byte("garbage"))

	if result != nil {
		t.Fail()
	}

	// #77 block: missing genesis block hash
	result = b.FindByBlockHash(genesisBlock.BlockHash)
	if !block.Equal(genesisBlock, result) {
		t.Fail()
	}
}

func TestBlockchain_FindByParenthash(t *testing.T) {
	b, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	target := b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))

	result := b.FindByParentHash(target.ParentHash)

	if !block.Equal(target, result) {
		t.Fail()
	}

	result = b.FindByParentHash([]byte("garbage"))

	if result != nil {
		t.Fail()
	}
}

func TestBlockchain_FindByTimestamp(t *testing.T) {
	b, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	target := b.AddSHA512([]byte("NewSTring"))
	time.Sleep(100 * time.Millisecond)
	b.AddSHA512([]byte("NewSTring2"))

	result := b.FindByTimestamp(target.Timestamp)

	if !block.Equal(target, result) {
		log.Println("target", target)
		log.Println("result", result)
		t.Error("Failed to find target with timestamp:", target.Timestamp, "resulting timestamp: ", result.Timestamp)
	}

	result = b.FindByTimestamp(1234)

	if result != nil {
		t.Fail()
	}
}

func TestBlockchain_UpdateChain(t *testing.T) {
	log.Println("Testing UpdateChain")
	b, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))
	c := b
	c.AddSHA512([]byte("new"))

	cExpectedLength := c.Length()
	updatedChain, err := UpdateChain(b, c)
	if err != nil {
		t.Error(err)
	}

	if !Equal(updatedChain, c) {
		t.Error("faild to return equal chains")
	}

	updatedChain.AddSHA512([]byte("asdf"))

	if c.Length() != cExpectedLength {
		t.Error("update chain did not return a copy")
	}

	d, err := New(genesisBlock)
	if err != nil {
		t.Error(err)
	}
	d.AddSHA512([]byte("invalid"))

	updatedChain, err = UpdateChain(c, d)
	if err == nil {
		t.Error(err)
	}
}
