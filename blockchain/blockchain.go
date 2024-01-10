package blockchain

import (
	"fmt"

	"github.com/LucioSchiavoni/blockchain/block"
)

type Blockchain struct {
	transactionPool []string
	chain           []*block.Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *block.Block {
	b := block.NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("Chain %d \n", i)
		block.Print()
		fmt.Println()

	}
}
