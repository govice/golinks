package blockchain

import (
	"testing"
)

func TestValidChain(t *testing.T) {
	blkchain := New()
	blkchain.Add("NewSTring")
	blkchain.Add("NewSTring")
	err := blkchain.Validate()
	if err != nil {
		t.Error("Could not validate blockchain")
	}
}

func TestGetValidChain(t *testing.T) {

	blkchain := New()
	blkchain.Add("NewSTring")
	blkchain.Add("NewSTring")
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

func TestBinaryConverter(t *testing.T) {
	blkchain := New()
	blkchain.Add("NewSTring")
	blkchain.Add("NewSTring")
	_ = blkchain.Validate()

	bin := blkchain.encodeChain()

	var out Blockchain
	err := out.decodeChain(bin)
	if err != nil {
		t.Error("Invalid decode Blockchain", err)
	}
}
