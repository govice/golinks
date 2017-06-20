package main

import (
	"github.com/LaughingCabbage/goLinks/types/blockchain"

)


func main() {
	blkchain := blockchain.New()
	blkchain.Add("NewSTring")
	blkchain.Add("NewSTring")
	blkchain.Print()

}
