package block

import (
	"log"
	"testing"
)

func TestEqual(t *testing.T) {

	//Test equal blocks
	blkA := New(0, []byte("data"), nil)
	blkB := blkA
	log.Println("Testing Block Equality")
	if !Equal(blkA, blkB) {
		t.Error("Block equivilents do not match")
	}

	//Test unequal blocks
	log.Println("Testing Block Inequality")
	blkC := New(1, []byte("str"), nil)

	if Equal(blkA, blkC) {
		t.Error("Unequivilent blocks returning as matching")
	}

}
