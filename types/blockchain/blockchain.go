package blockchain

import(
	"container/list"
	"github.com/LaughingCabbage/goLinks/types/block"
	"math/cmplx"
)

var index int
var Blockchain *list.List

func New() {
	//create a new blockchain starting with a genesis block
	index = 0
	Blockchain = list.New()
	blk := block.New(index, "GENESIS DATA", nil)
	Blockchain.PushBack(blk)
}

func Add(data string) {
	for i := Blockchain.Front(); i != index; i++{

	}
	index++
	blk := block.New(index, data, Blockchain.)
	block.Index = index
	Blockchain.PushBack(block)
}
