package blockchain

import(
	"github.com/LaughingCabbage/goLinks/types/block"

	"fmt"
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
