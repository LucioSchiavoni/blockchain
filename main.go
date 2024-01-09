package main

import (
	"fmt"

	"github.com/LucioSchiavoni/blockchain/block"
)

func main() {

	fmt.Println("Hello block")
	b := block.NewBlock(0, "init Hash")
	b.Print()

}
