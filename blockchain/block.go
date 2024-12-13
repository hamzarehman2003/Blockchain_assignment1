package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int       // Position of the block in the chain
	Timestamp    time.Time // Time when the block was created
	Data         string    // Data stored in the block
	PreviousHash string    // Hash of the previous block
	Hash         string    // Current block's hash
}

// Function to compute the hash of the block
func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp.String() + b.Data + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// Function to create a new block
func NewBlock(index int, data string, previousHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now(),
		Data:         data,
		PreviousHash: previousHash,
	}
	block.Hash = block.CalculateHash()
	return block
}

func main() {
	// Example: Create a genesis block
	genesisBlock := NewBlock(0, "Genesis Block", "")
	
	// Example: Add another block
	newBlock := NewBlock(1, "Second Block", genesisBlock.Hash)

	// Print the blocks
	println("Genesis Block:")
	println("Index:", genesisBlock.Index)
	println("Timestamp:", genesisBlock.Timestamp.String())
	println("Data:", genesisBlock.Data)
	println("Previous Hash:", genesisBlock.PreviousHash)
	println("Hash:", genesisBlock.Hash)

	println("\nNew Block:")
	println("Index:", newBlock.Index)
	println("Timestamp:", newBlock.Timestamp.String())
	println("Data:", newBlock.Data)
	println("Previous Hash:", newBlock.PreviousHash)
	println("Hash:", newBlock.Hash)
}
