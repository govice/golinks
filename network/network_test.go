package network

import (
	"bytes"
	"encoding/gob"
	"log"
	"testing"

	"github.com/LaughingCabbage/goLinks/types/block"
)

func TestGenerateKeyPair(t *testing.T) {
	_, err := GenerateKeyPair()
	if err != nil {
		t.Error("faield to generate key pair", err)
	}
}

func TestBinaryConverters(t *testing.T) {
	var network bytes.Buffer

	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	blk := block.New(0, []byte("data"), nil)
	err := enc.Encode(blk)
	if err != nil {
		t.Error("Binary Encoding Failed", err)
	}

	var b block.Block
	err = dec.Decode(&b)
	if err != nil {
		t.Error("Binary Decoding Failed ", err)
	}
	if blk.Index != b.Index {
		t.Error("Block indexes do not corrolate")
	}
	if blk.Timestamp != b.Timestamp {
		log.Print(blk.Timestamp)
		log.Print(b.Timestamp)
		t.Error("block timestamps do not corrolate")
	}
	if bytes.Equal(blk.Data, b.Data) {
		t.Error("Block data does not corrolate")
	}
	if !bytes.Equal(blk.Parenthash, b.Parenthash) {
		t.Error("BLock parent hashes do not corrolate")
	}
	if !bytes.Equal(blk.Blockhash, b.Blockhash) {
		t.Error("Block hashes do not corrolate")
	}
}
