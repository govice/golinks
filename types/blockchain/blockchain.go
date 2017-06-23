package blockchain

import(
	"github.com/LaughingCabbage/goLinks/types/block"

	"fmt"
	"errors"
	"bytes"
	"encoding/gob"
	"log"
)

type Blockchain []block.Block


func New() Blockchain {
	var blkchain Blockchain
	//create genesis block and append it as root to blockchain
	blk := block.New(0, "GENESIS DATA", nil)
	blkchain = append(blkchain, blk)
	return blkchain
}

func (blockchain *Blockchain) Add(data string){
	blk := block.New(len(*blockchain), data, (*blockchain)[len(*blockchain)-1].Blockhash)
	*blockchain = append(*blockchain, blk)
}

func (blockchain Blockchain) Print(){
	for i := 0; i < len(blockchain); i++{
		fmt.Println(i)
		fmt.Println(blockchain[i])
	}
}

func (blockchain Blockchain) Validate() error{
	if(len(blockchain) < 2){
		return errors.New("invalid attempt to validate genesis block")
	}
	for i := 1; i < len(blockchain); i++{
		err := block.Validate(blockchain[i-1], blockchain[i])
		if err != nil{
			return errors.New("blockchain is invalid")
		}
	}
	return nil
}

func (blockchain Blockchain) EncodeChain() []byte{
	buffer := bytes.Buffer{}
	chainGob := gob.NewEncoder(&buffer)
	err := chainGob.Encode(blockchain)
	if err != nil{
		log.Fatal("failed to encode blockchain")
	}
	return buffer.Bytes()

}