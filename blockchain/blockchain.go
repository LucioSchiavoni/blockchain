package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/s2a-go/retry"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlock(nonce int, previusHash [32]byte, transactions []*Transaction) *Block {

	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previusHash
	b.transactions = transactions
	return b
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func BlockHash() {
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json: "nonce"`
		PreviousHash [32]byte       `json: "previous_hash"`
		Transactions []*Transaction `json: "transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("Timestamp:  %d\n", b.timestamp)
	fmt.Printf("Nonce:  %d\n", b.nonce)
	fmt.Printf("Previus Hash:  %x\n", b.previousHash)
	fmt.Printf("Transactions:  %s\n", b.transactions)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction{
	transactions := make([]*Transaction, 0)
	for _, t:= range bc.transactionPool{
		transactions = append(transactions, 
		NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value)
		)
	}
	return transactions 
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}

}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_address %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient address %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json: "sender_blockchain_address"`
		Recipient string  `json: "recipient_blockchain_address"`
		Value     float32 `json: "value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
