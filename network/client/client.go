package client

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/LaughingCabbage/goLinks/network"
)

type Client struct {
	serverAddress string
	serverPort    string
	serverConn    net.Conn
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

func (client *Client) Connect() error {
	var err error
	client.serverConn, err = net.Dial("tcp", lookupServer())
	if err != nil {
		return errors.New("failed to dial server")
	}
	return nil

}

func lookupServer() string {
	return "localhost:8080"
}

func (client Client) Message(message string) {
	if client.serverConn == nil {
		log.Fatal("client's server connection not set")
		return
	}
	fmt.Fprintf(client.serverConn, message+"\n")

}

func (client Client) Listen() {
	if client.serverConn == nil {
		return
	}
	message, _ := bufio.NewReader(client.serverConn).ReadString('\n')
	fmt.Print("Server respone: " + message)
}

func (client Client) Print() {
	fmt.Print(client)
}

func (client *Client) newKey() error {
	var err error
	client.privateKey, err = network.GenerateKeyPair()
	if err != nil {
		log.Fatalf("failed to assign key to client ", err)
	}
	return err
}
