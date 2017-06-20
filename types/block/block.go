package block

import (
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



func New(index int, data string, parent []uint8) Block {
	blk := Block{index, time.Now(), data, nil, nil}
	//handle genesis block case
	if index == 0 {
		parenthash := sha512.New()
		parenthash.Write([]byte(blk.Data))
		blk.Parenthash = parenthash.Sum(nil)
	}else{
		blk.Parenthash = parent
	}
	blkhash := sha512.New()
	blkhash.Write(EncodeGob(blk))
	blk.Blockhash = blkhash.Sum(nil)
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
