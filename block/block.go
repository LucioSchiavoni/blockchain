package block

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previusHash string) *Block {

	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previusHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp   %d\n", b.timestamp) //%d para int y %s para string
	fmt.Printf("nonce   %d\n", b.nonce)
	fmt.Printf("previous Hash   %s\n", b.previousHash)
	fmt.Printf("transactions   %s\n", b.transactions)

}
