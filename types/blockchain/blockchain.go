package blockchain

import(
	"container/list"
	"github.com/LaughingCabbage/goLinks/types/block"
)


func New() *list.List {
	//create a new blockchain starting with a genesis block
	blkchain := list.New()
	blk := block.New(0, nil)
	blkchain.PushBack(blk)
	return blkchain

}

