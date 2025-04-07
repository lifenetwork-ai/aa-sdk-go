package aasdk

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	MessagePrefix = "\x19Ethereum Signed Message:\n"
)

// SignMessage signs a message with the provided private key.
func SignMessage(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	prefixedMessage := fmt.Sprintf("%s%d", MessagePrefix, len(message))
	bytes := append([]byte(prefixedMessage), message...)
	hash := crypto.Keccak256Hash([]byte(bytes))
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}
	signature[crypto.RecoveryIDOffset] += 27
	return signature, nil
}
