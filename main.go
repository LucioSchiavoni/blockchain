package main

import (
	"github.com/LucioSchiavoni/blockchain/blockchain"
)

func main() {

	blockGenerate := blockchain.NewBlockchain()
	blockGenerate.Print()

	blockchain.NewBlockchain().AddTransaction("A", "B", 1.0)

	previousHash := blockGenerate.LastBlock().Hash()
	blockGenerate.CreateBlock(5, previousHash)
	blockGenerate.Print()

	previousHash = blockGenerate.LastBlock().Hash()
	blockGenerate.CreateBlock(2, previousHash)
	blockGenerate.Print()
}
