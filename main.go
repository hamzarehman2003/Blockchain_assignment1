// Move the startAPI function to main.go
package main

import (
	"Blockchain_assignment1/blockchain"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var otherNodes = []string{
	"http://node2:8002",
	"http://node3:8003",
	"http://node4:8004",
}

var mempool []blockchain.Transaction
var blockchainChain []blockchain.Block
var mutex = &sync.Mutex{} // To handle concurrent access

// Initialize the blockchain
func initBlockchain() {
	genesisBlock := blockchain.NewBlock(0, "Genesis Block", "")
	blockchainChain = append(blockchainChain, genesisBlock)
}

// API to receive a transaction
func handleTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var tx blockchain.Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	mempool = append(mempool, tx)
	mutex.Unlock()
	fmt.Fprintln(w, "Transaction added to mempool")
}

// API to mine a block
func mineBlock(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	if len(mempool) == 0 {
		http.Error(w, "No transactions to mine", http.StatusBadRequest)
		return
	}

	lastBlock := blockchainChain[len(blockchainChain)-1]
	newBlock := blockchain.NewBlock(lastBlock.Index+1, fmt.Sprintf("%v", mempool), lastBlock.Hash)
	blockchainChain = append(blockchainChain, newBlock)
	mempool = []blockchain.Transaction{}

	// Broadcast the new block
	broadcastBlock(newBlock)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBlock)
}

// API to return the blockchain
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchainChain)
}

// API to accept a mined block
func handleBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var newBlock blockchain.Block
	err := json.NewDecoder(r.Body).Decode(&newBlock)
	if err != nil {
		http.Error(w, "Invalid block data", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if newBlock.Index == blockchainChain[len(blockchainChain)-1].Index+1 &&
		newBlock.PreviousHash == blockchainChain[len(blockchainChain)-1].Hash {
		blockchainChain = append(blockchainChain, newBlock)
	} else {
		fmt.Println("Invalid block received")
	}
	w.WriteHeader(http.StatusOK)
}

// Broadcast a mined block to all nodes
func broadcastBlock(block blockchain.Block) {
	for _, node := range otherNodes {
		go func(nodeURL string) {
			data, _ := json.Marshal(block)
			resp, err := http.Post(fmt.Sprintf("%s/block", nodeURL), "application/json", io.NopCloser(bytes.NewReader(data)))
			if err != nil {
				fmt.Printf("Error broadcasting block to %s: %v\n", nodeURL, err)
				return
			}
			resp.Body.Close()
		}(node)
	}
}

// Main API handler
func startAPI(port string) {
	http.HandleFunc("/transaction", handleTransaction)
	http.HandleFunc("/mine", mineBlock)
	http.HandleFunc("/blockchain", getBlockchain)
	http.HandleFunc("/block", handleBlock)

	fmt.Printf("Node running on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// Main function
func main() {
	port := "8001" // Port for this node
	fmt.Println("Starting node...")
	initBlockchain()
	startAPI(port)
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
