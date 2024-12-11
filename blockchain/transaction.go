package blockchain

import (
	"errors"
	"time"
)

// Transaction represents a blockchain transaction
type Transaction struct {
	ID         string    // Unique identifier for the transaction
	Sender     string    // Address of the sender
	Receiver   string    // Address of the receiver
	Amount     float64   // Amount to be transferred
	Timestamp  time.Time // Timestamp of the transaction
	AlgoCid    string
	DatasetCid string
	Algo       string 	 // Algo name
	Signature  []string
	ResultHash string
}

// GetID returns the ID of the transaction
func (t *Transaction) GetID() string {
	return t.ID
}

// GetSender returns the sender address of the transaction
func (t *Transaction) GetSender() string {
	return t.Sender
}

// GetReceiver returns the receiver address of the transaction
func (t *Transaction) GetReceiver() string {
	return t.Receiver
}

// GetAmount returns the amount of the transaction
func (t *Transaction) GetAmount() float64 {
	return t.Amount
}

// GetTimestamp returns the timestamp of the transaction
func (t *Transaction) GetTimestamp() time.Time {
	return t.Timestamp
}

// GetAlgoCid returns the Algo CID of the transaction
func (t *Transaction) GetAlgoCid() string {
	return t.AlgoCid
}

// GetDatasetCid returns the Dataset CID of the transaction
func (t *Transaction) GetDatasetCid() string {
	return t.DatasetCid
}

// GetAlgo returns the algorithm name of the transaction
func (t *Transaction) GetAlgo() string {
	return t.Algo
}

// GetSignature returns the signature of the transaction
func (t *Transaction) GetSignature() []string {
	return t.Signature
}

// GetResultHash returns the result hash of the transaction
func (t *Transaction) GetResultHash() string {
	return t.ResultHash
}

// CreateTransaction initializes and returns a new Transaction
func CreateTransaction(
	id string,
	sender string,
	receiver string,
	amount float64,
	algoCid string,
	datasetCid string,
	algo string,
	signature []string,
	resultHash string,
) (*Transaction, error) {
	// Validation: Ensure required fields are not empty
	if id == "" || sender == "" || receiver == "" || amount <= 0 {
		return nil, errors.New("invalid transaction parameters: ID, sender, receiver, and amount must be valid")
	}

	// Create the transaction
	tx := &Transaction{
		ID:         id,
		Sender:     sender,
		Receiver:   receiver,
		Amount:     amount,
		Timestamp:  time.Now(),
		AlgoCid:    algoCid,
		DatasetCid: datasetCid,
		Algo:       algo,
		Signature:  signature,
		ResultHash: resultHash,
	}

	return tx, nil

}
