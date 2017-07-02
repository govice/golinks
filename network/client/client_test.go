package client

import (
	"testing"
)

func TestNewKey(t *testing.T) {
	var c Client
	err := c.newKey()
	if err != nil {
		t.Error("test failed to generate key pair")
	}
}
