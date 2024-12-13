package main

import (
	"fmt"
	"time"

	"Blockchain_assignment1/blockchain"

)

func main() {
	// Create the genesis block
	genesisBlock := blockchain.NewBlock(0, "Genesis Block", "")

	// Add a second block
	secondBlock := blockchain.NewBlock(1, "Second Block", genesisBlock.Hash)

	// Print the blocks
	fmt.Println("Genesis Block:")
	fmt.Println("Index:", genesisBlock.Index)
	fmt.Println("Timestamp:", genesisBlock.Timestamp.Format(time.RFC3339))
	fmt.Println("Data:", genesisBlock.Data)
	fmt.Println("Previous Hash:", genesisBlock.PreviousHash)
	fmt.Println("Hash:", genesisBlock.Hash)

	fmt.Println("\nSecond Block:")
	fmt.Println("Index:", secondBlock.Index)
	fmt.Println("Timestamp:", secondBlock.Timestamp.Format(time.RFC3339))
	fmt.Println("Data:", secondBlock.Data)
	fmt.Println("Previous Hash:", secondBlock.PreviousHash)
	fmt.Println("Hash:", secondBlock.Hash)
}
