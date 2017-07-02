package network

import "testing"

func TestGenerateKeyPair(t *testing.T) {
	_, err := GenerateKeyPair()
	if err != nil {
		t.Error("faield to generate key pair", err)
	}
}
