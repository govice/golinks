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
	"log"
	"os"
	"testing"
	"time"

	"github.com/govice/golinks/block"
)

var genesisBlock = block.NewSHA512(0, []byte("GENESIS"), nil)

//TestGetValidChain creates two different blockchains of different sizes and attempts to validate the chain.
func TestBlockchain_Validate(t *testing.T) {
	log.Println("Testing GetValidChain")
	blkchain := New(genesisBlock)
	blkchain.AddSHA512([]byte("NewSTring"))
	blkchain.AddSHA512([]byte("NewSTring"))
	err := blkchain.Validate()
	if err != nil {
		t.Error("Failed to Valiate Blockchain")
	}
	blkchain2 := New(genesisBlock)
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
	chainA := New(genesisBlock)
	var chainB Blockchain

	//Test with equal blocks
	chainA.AddSHA512([]byte("NewSTring"))
	chainA.AddSHA512([]byte("NewSTring"))

	chainB = chainA

	//Test equality with two equal blockchains
	if !Equal(chainA, chainB) {
		t.Error("equal chains are not equal")
	}

	//Test equality with additional block
	chainB.AddSHA512([]byte("data"))
	if Equal(chainA, chainB) {
		t.Error("unequal chains are testing as equal")
	}

}

func TestBlockchain_InputOutput(t *testing.T) {
	log.Println("Testing I/O")
	//Test saving to file
	blkchain := New(genesisBlock)
	blkchain.AddSHA512([]byte("NewSTring"))
	blkchain.AddSHA512([]byte("NewSTring2"))

	err := blkchain.Save("testfile")
	if err != nil {
		t.Error("failed to save blockchain ", err)
	}

	//Test loading from file
	blkchainB := Blockchain{}
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
	b := New(genesisBlock)
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
	b := New(genesisBlock)
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

func TestBlockchain_UpdateChain(t *testing.T) {
	log.Println("Testing UpdateChain")
	b := New(genesisBlock)
	b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))
	c := b
	c.AddSHA512([]byte("new"))
	err := b.UpdateChain(c)
	if err != nil {
		t.Error(err)
	}
	if !Equal(c, b) {
		t.Error("Failed Update Chain")
	}

	d := New(genesisBlock)
	d.AddSHA512([]byte("invalid"))
	err = c.UpdateChain(d)
	if err == nil {
		t.Error(err)
	}
}

func TestBlockchain_FindByHash(t *testing.T) {
	b := New(genesisBlock)
	target := b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))

	result := b.FindByBlockHash(target.Blockhash())
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
}

func TestBlockchain_FindByParenthash(t *testing.T) {
	b := New(genesisBlock)
	target := b.AddSHA512([]byte("NewSTring"))
	b.AddSHA512([]byte("NewSTring2"))

	result := b.FindByParentHash(target.Parenthash())

	if !block.Equal(target, result) {
		t.Fail()
	}

	result = b.FindByParentHash([]byte("garbage"))

	if result != nil {
		t.Fail()
	}
}

func TestBlockchain_FindByTimestamp(t *testing.T) {
	b := New(genesisBlock)
	target := b.AddSHA512([]byte("NewSTring"))
	time.Sleep(100 * time.Millisecond)
	b.AddSHA512([]byte("NewSTring2"))

	result := b.FindByTimestamp(target.Timestamp())

	if !block.Equal(target, result) {
		log.Println("target", target)
		log.Println("result", result)
		t.Error("Failed to find target with timestamp:", target.Timestamp(), "resulting timestamp: ", result.Timestamp())
	}

	result = b.FindByTimestamp(1234)

	if result != nil {
		t.Fail()
	}
}
