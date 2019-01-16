package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/laughingcabbage/golinks/blockchain"
)

func checkEnv() {
	if govicePath == "" {
		log.Fatal("govicePath env variable undefined")
	}
	log.Println("GOVICE", govicePath)
}

func loadBlockchain() {
	viceChainPath = govicePath + string(os.PathSeparator) + chainName + ".dat"
	log.Println("Checking for existing chain", viceChainPath)
	if _, err := os.Stat(viceChainPath); os.IsNotExist(err) {
		chain = blockchain.New()
		log.Println("Generated new blockchain")
		chain.Print()
		if err := chain.Save(chainName); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := chain.Load(chainName); err != nil {
			log.Fatal(err)
		}
		log.Println("Loaded existing blockchain")
	}
}

func loadInfo() {
	infoFile, err := os.Open("info.json")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(infoFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := infoFile.Close(); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &info); err != nil {
		log.Fatal(err)
	}

}
