package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Connect to the IPFS daemon API server
	sh := shell.NewShell("localhost:5001") // Ensure your IPFS daemon is running

	// 1. Upload a file to IPFS
	filePath := "data.txt"
	cid, err := uploadFile(sh, filePath)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}
	fmt.Printf("File uploaded successfully! CID: %s\n", cid)

	// 2. Download the file from IPFS
	outputPath := "downloaded_data.txt"
	err = downloadFile(sh, cid, outputPath)
	if err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}
	fmt.Printf("File downloaded successfully to %s\n", outputPath)
}

// uploadFile uploads a file to IPFS and returns its CID
func uploadFile(sh *shell.Shell, filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to upload file to IPFS: %w", err)
	}
	return cid, nil
}

// downloadFile downloads a file from IPFS using its CID
func downloadFile(sh *shell.Shell, cid, outputPath string) error {
	data, err := sh.Cat(cid)
	if err != nil {
		return fmt.Errorf("failed to download file from IPFS: %w", err)
	}

	// Save the data to a file
	content, err := ioutil.ReadAll(data)
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}

	err = ioutil.WriteFile(outputPath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
