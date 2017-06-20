package block

import (
	"fmt"
	"time"
	"crypto/sha512"
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Index      int
	Timestamp  time.Time
	Data       string
	Parenthash []uint8
	Blockhash  []uint8
}

func New() Block {
	blk := Block{-1, time.Now(), "TEST DATAA", nil, nil}
	hash := sha512.New()
	hash.Write([]byte(blk.Data))
	blk.Parenthash = hash.Sum(nil)
	hash = sha512.New()
	hash.Write(EncodeGob(blk))
	blk.Blockhash = hash.Sum(nil)
	fmt.Println(blk)
	return blk
}

func EncodeGob(block Block) []byte{
	//we need to serialize everything but the current block hash
	//create a buffer, encode gob to buffer, return serial byte array
	buffer := bytes.Buffer{}
	blockgob := gob.NewEncoder(&buffer)
	err := blockgob.Encode(Block{block.Index, block.Timestamp, block.Data, block.Parenthash, nil})
	if err != nil {
		log.Fatal("failed to encode block into Gob")
	}
	return buffer.Bytes()
}
