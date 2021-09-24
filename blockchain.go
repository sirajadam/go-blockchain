package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block for storing information
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

// SetHash creates
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	
	b.Hash = hash[:]
}

// NewBlock creates new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash,[]byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates the first block of the Blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}


// Blockchain a database with certain structure
type BlockChain struct {
	blocks []*Block
}

// AddBlock adds new block to the Blockchain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	genesisBlock := NewGenesisBlock()
	return &BlockChain{[]*Block{genesisBlock}}
}





func main() {

	bc := NewBlockChain()

	bc.AddBlock("Send 1700 AKT to Adam")
	bc.AddBlock("Send 300 AKT to anon")

	for _, block := range bc.blocks {
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("Previous Hash: %x\n",block.PrevBlockHash)
	}




}