package blockchain

import (
	"testing"
)

func TestValidChain(t *testing.T) {
	blkchain := New()
	blkchain.Add("NewSTring")
	blkchain.Add("NewSTring")
	blkchain.Print()
	err := blkchain.Validate()
	if err != nil {
		t.Error("Could not validate blockchain")
	}
}

func TestGetValidChain(t *testing.T) {

	blkchain := New()
	blkchain.Add("NewSTring")
	blkchain.Add("NewSTring")
	blkchain.Print()
	_ = blkchain.Validate()

	blkchain2 := New()
	blkchain2.Add("chain2str")
	blkchain2.Add("chain2str")
	blkchain2.Add("data")
	blkchain2.Validate()

	validChain := GetValidChain(blkchain, blkchain2)

	if validChain[0].Data != blkchain2[0].Data {
		t.Error("Valid block was not returned")
	}
}
