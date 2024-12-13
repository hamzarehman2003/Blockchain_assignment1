package ipfs

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

// IPFSClient wraps the IPFS shell for reusability
type IPFSClient struct {
	shell *shell.Shell
}

// NewIPFSClient creates a new IPFS client connected to the given IPFS daemon API server
func NewIPFSClient(apiAddress string) (*IPFSClient, error) {
	sh := shell.NewShell(apiAddress)
	if !sh.IsUp() {
		return nil, fmt.Errorf("failed to connect to IPFS daemon at %s", apiAddress)
	}
	log.Printf("Connected to IPFS daemon at %s", apiAddress)
	return &IPFSClient{shell: sh}, nil
}

// UploadFile uploads a file to IPFS and returns its CID
func (client *IPFSClient) UploadFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	cid, err := client.shell.Add(bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to upload file to IPFS: %w", err)
	}
	return cid, nil
}

// DownloadFile downloads a file from IPFS using its CID and saves it to the given output path
func (client *IPFSClient) DownloadFile(cid, outputPath string) error {
	data, err := client.shell.Cat(cid)
	if err != nil {
		return fmt.Errorf("failed to download file from IPFS: %w", err)
	}
	defer data.Close()

	content, err := io.ReadAll(data)
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}

	err = ioutil.WriteFile(outputPath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
