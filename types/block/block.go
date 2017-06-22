package block

import (
	"time"
	"crypto/sha512"
	"bytes"
	"encoding/gob"
	"log"
	"errors"
)

type Block struct {
	Index      int
	Timestamp  time.Time
	Data       string
	Parenthash []byte
	Blockhash  []byte
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
	blk.Blockhash = Hash(blk)
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

//calculate and return the hash for a block
func Hash(block Block) []uint8{
	blkhash := sha512.New()
	blkhash.Write(EncodeGob(block))
	return blkhash.Sum(nil)
}

//need to check index, and hashes
func Validate(prev, current Block) error{
	if prev.Index+1 != current.Index {
		return errors.New("block indexes do not corrilate")
	}
	if !bytes.Equal(prev.Blockhash, current.Parenthash){
		return errors.New("block parent child hashes do not corrilate")
	}
	if !bytes.Equal(Hash(current), current.Blockhash) {
		return errors.New("current block's hash is not valid")
	}
	return nil
	
}
