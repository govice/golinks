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

package blockchain

import (
	"log"
	"os"
	"testing"
)

//TestGetValidChain creates two different blockchains of different sizes and attempts to validate the chain.
func TestBlockchain_Validate(t *testing.T) {
	log.Println("Testing GetValidChain")
	blkchain := New()
	blkchain.Add([]byte("NewSTring"))
	blkchain.Add([]byte("NewSTring"))
	err := blkchain.Validate()
	if err != nil {
		t.Error("Failed to Valiate Blockchain")
	}
	blkchain2 := New()
	blkchain2.Add([]byte("chain2str"))
	blkchain2.Add([]byte("chain2str"))
	blkchain2.Add([]byte("data"))
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
	chainA := New()
	var chainB Blockchain

	//Test with equal blocks
	chainA.Add([]byte("NewSTring"))
	chainA.Add([]byte("NewSTring"))

	chainB = chainA

	//Test equality with two equal blockchains
	if !Equal(chainA, chainB) {
		t.Error("equal chains are not equal")
	}

	//Test equality with additional block
	chainB.Add([]byte("data"))
	if Equal(chainA, chainB) {
		t.Error("unequal chains are testing as equal")
	}

}

func TestBlockchain_InputOutput(t *testing.T) {
	log.Println("Testing I/O")
	//Test saving to file
	blkchain := New()
	blkchain.Add([]byte("NewSTring"))
	blkchain.Add([]byte("NewSTring2"))
	err := blkchain.Save("testfile")
	if err != nil {
		t.Error("failed to save blockchain ", err)
	}

	//Test loading from file
	var blkchainB Blockchain
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

func TestBlockchain_GetGCI(t *testing.T) {
	log.Println("Testing GetGCI")
	b := New()
	b.Add([]byte("NewSTring"))
	b.Add([]byte("NewSTring2"))
	c := b
	c.Add([]byte("new"))
	if b.GetGCI(c) != 3 {
		t.Errorf("Invalid GCI of %v should be 3", len(b))
	}

}

func TestBlockchain_UpdateChain(t *testing.T) {
	log.Println("Testing UpdateChain")
	b := New()
	b.Add([]byte("NewSTring"))
	b.Add([]byte("NewSTring2"))
	c := b
	c.Add([]byte("new"))
	err := b.UpdateChain(c)
	if err != nil {
		t.Error(err)
	}
	if !Equal(c, b) {
		t.Error("Failed Update Chain")
	}

	d := New()
	d.Add([]byte("invalid"))
	err = c.UpdateChain(d)
	if err == nil {
		t.Error(err)
	}

}
