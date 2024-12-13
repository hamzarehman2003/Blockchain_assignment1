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

// CalculateHash computes the hash for the block
func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp.String() + b.Data + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// NewBlock creates a new block with the given data
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
