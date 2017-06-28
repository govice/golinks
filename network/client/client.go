package client

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

type Client struct {
	serverAddress string
	serverPort    string
	privateKey    *rsa.PrivateKey
	publicKey     *rsa.PublicKey
	serverKey     *rsa.PublicKey
}

//New creates a new client and provides the server address.
func New(serverAddr, serverPort string) Client {
	var client Client
	client.serverAddress = serverAddr
	client.serverPort = serverPort
	var err error
	client.privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("failed to generate RSA private key ", err)
	}
	return client
	
	//entropy := []byte("laughingcabbage%$#@")
	//label := []byte("message")

	/*
		rng := rand.Reader

		ciphertext, err := rsa.EncryptOAEP(sha512.New(), rng, &client.publicKey, entropy, label)
		if err != nil {
			//		log.Fatal("Error in encryption ", err)
		}
		fmt.Println("ciphertext ", ciphertext)
	*/
}
