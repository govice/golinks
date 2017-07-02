package network

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func GenerateKeyPair() (*rsa.PrivateKey, error) {
	keySize := 1024
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		log.Fatalf("failed to generate private key ", err)
	}
	return privateKey, err
}
