package main

import(
	"time"
	"fmt"
	"github.com/LaughingCabbage/goLinks/types/block"

)

func main() {
	t0 := time.Now()
	fmt.Println(t0.String())

	rootBlock := block.New()

	fmt.Println(rootBlock.Parenthash)

}