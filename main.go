package main

import (
	"fmt"

	"github.com/LucioSchiavoni/blockchain/block"
	"github.com/LucioSchiavoni/blockchain/blockchain"
)

func main() {

	b := block.NewBlock(0, "init Hash")
	b.Print()
	fmt.Println("--------------")
	bc := blockchain.NewBlockchain().CreateBlock(5, "hash 1")
	bc.Print()
	fmt.Println("--------------")
	bc1 := blockchain.NewBlockchain().CreateBlock(2, "hash 2")
	bc1.Print()
	fmt.Println("--------------")
}
