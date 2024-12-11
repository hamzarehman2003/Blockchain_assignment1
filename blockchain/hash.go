package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

func HashAnything(input interface{}) string {
	// Convert the input to a string representation
	var data string

	switch v := input.(type) {
	case string:
		data = v
	case []byte:
		data = string(v)
	case int, int64, float64, bool:
		data = fmt.Sprintf("%v", v)
	default:
		// Use reflection for more complex types
		data = fmt.Sprintf("%v", reflect.ValueOf(input))
	}

	// Create a SHA-256 hash
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
