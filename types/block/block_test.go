package block

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestBinaryConverters(t *testing.T) {
	var network bytes.Buffer

	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	blk := New(0, "data", nil)
	err := enc.Encode(blk)
	if err != nil {
		t.Error("Binary Encoding Failed", err)
	}

	var b Block
	err = dec.Decode(&b)
	if err != nil {
		t.Error("Binary Decoding Failed ", err)
	}
	if blk.Index != b.Index {
		t.Error("Block indexes do not corrolate")
	}
	if blk.Timestamp != b.Timestamp {
		t.Error("block timestamps do not corrolate")
	}
	if blk.Data != b.Data {
		t.Error("Block data does not corrolate")
	}
	if !bytes.Equal(blk.Parenthash, b.Parenthash) {
		t.Error("BLock parent hashes do not corrolate")
	}
	if !bytes.Equal(blk.Blockhash, b.Blockhash) {
		t.Error("Block hashes do not corrolate")
	}
}

func TestEqual(t *testing.T) {

	//Test equal blocks
	blkA := New(0, "data", nil)
	blkB := blkA

	if !Equal(blkA, blkB) {
		t.Error("Block equivilents do not match")
	}

	//Test unequal blocks
	blkC := New(1, "str", nil)

	if Equal(blkA, blkC) {
		t.Error("Unequivilent blocks returning as matching")
	}
}
