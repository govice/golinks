package block

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {

	//Test equal blocks
	blkA := NewSHA512(0, []byte("data"), nil)
	blkB := blkA
	log.Println("Testing Block Equality")
	if !bytes.Equal(blkA.BlockHash, blkB.BlockHash) {
		t.Fail()
	}

	//Test unequal blocks
	log.Println("Testing Block Inequality")
	blkC := NewSHA512(1, []byte("str"), nil)

	if bytes.Equal(blkA.BlockHash, blkC.BlockHash) {
		t.Error("Unequivilent blocks returning as matching")
	}

	target := blkB

	if !Equal(blkB, target) {
		t.Fail()
	}

	if Equal(blkB, blkC) {
		t.Fail()
	}
}

func TestJSON(t *testing.T) {
	log.Println("Testing block JSON")
	blkA := NewSHA512(0, []byte("GENESIS"), nil)
	blkB := NewSHA512(1, []byte("data"), blkA.BlockHash)

	jsonBytes, err := json.Marshal(blkB)
	if err != nil {
		t.Error(err)
	}

	blockOut := &Block{}
	if err := json.Unmarshal(jsonBytes, blockOut); err != nil {
		t.Error(err)
	}

	if blkB.Index != blockOut.Index || blkB.Timestamp != blockOut.Timestamp {
		t.Fail()
	}

	if !bytes.Equal(blkB.ParentHash, blockOut.ParentHash) {
		t.Fail()
	}

	if !bytes.Equal(blkB.BlockHash, blockOut.BlockHash) {
		t.Fail()
	}

	if !bytes.Equal(blkB.Data, blockOut.Data) {
		t.Fail()
	}
}

// TODO better test. serialize omits the current block's hash when converting to json.
func TestSerialize(t *testing.T) {
	log.Println("Testing block serialize")
	blkA := NewSHA512(0, []byte("GENESIS"), nil)
	blkB := NewSHA512(1, []byte("data"), blkA.BlockHash)

	goldenBytes, err := json.Marshal(blkB)
	if err != nil {
		t.Error(err)
	}

	serializedBytes, err := blkB.Serialize()
	if err != nil {
		t.Error(err)
	}

	if bytes.Equal(goldenBytes, serializedBytes) {
		fmt.Printf("Golden: %x\n", goldenBytes)
		fmt.Printf("Serial: %x\n", serializedBytes)
		t.Fail()
	}
}

func TestNewGenesis(t *testing.T) {
	genesis := NewSHA512Genesis()
	timeZero := time.Time{}.UnixNano()
	if genesis.Timestamp != timeZero {
		t.Fail()
	}

	if genesis.Index != 0 {
		t.Fail()
	}

	if !bytes.Equal(genesis.Data, []byte("GENSIS BLOCK")) {
		t.Fail()
	}
}
