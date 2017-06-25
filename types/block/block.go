package block

import (
	"time"
	"crypto/sha512"
	"bytes"
	"errors"
	"fmt"
	"log"
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


//calculate and return the hash for a block
func Hash(block Block) []uint8{
	blkhash := sha512.New()
	buff, err := block.MarshalBinary()
	if err != nil {
		log.Fatal("Failed to hash block")
	}
	blkhash.Write(buff)
	return blkhash.Sum(nil)
}

//need to check index, and hashes
func Validate(prev, current Block) error{
	if prev.Index+1 != current.Index {
		return errors.New("block indexes do not correlate")
	}
	if !bytes.Equal(prev.Blockhash, current.Parenthash){
		return errors.New("block parent child hashes do not correlate")
	}
	if !bytes.Equal(Hash(current), current.Blockhash) {
		fmt.Println("Hash: ", Hash(current))
		fmt.Println("current: ", current.Blockhash)
		return errors.New("current block's hash is not valid")
	}
	return nil
}

func (block Block) MarshalBinary() ([]byte, error){
	var buffer bytes.Buffer
	fmt.Fprintln(&buffer, block.Index, block.Timestamp, block.Data, block.Parenthash)
	return buffer.Bytes(), nil
}

func (block *Block) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(buffer, &block.Index, &block.Timestamp, &block.Data, &block.Parenthash)
	return err
}
