package block

import(
	//"fmt"
	"time"
	//"hash"
	"crypto/sha512"
)
type Block struct {
	Index 		int
	Timestamp 	time.Time
	Data 		string
	Parenthash 	[]uint8
	Blockhash 	[]uint8

}

func New() Block{
	block := Block{-1, time.Now(), "TEST DATAA", nil, nil}
	hash := sha512.New()
	hash.Write([]byte(block.Data))
	block.Parenthash = hash.Sum(nil)
	return block
}

//func Computehash()
