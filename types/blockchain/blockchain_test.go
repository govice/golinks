package blockchain

import (
	"fmt"
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
	fmt.Println(blkchain)

	var out Blockchain
	err := out.decodeChain(bin)
	fmt.Println(out)
	if err != nil {
		t.Error("Invalid decode Blockchain", err)
	}

	/*
		original := []byte(blkchain)
		received := []byte(out)
		if !bytes.Equal(original, received) {
			t.Error("blockcahin deserialize does not match origin")
		}
	*/
}
