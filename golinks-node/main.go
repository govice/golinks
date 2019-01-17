package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/laughingcabbage/golinks/block"

	"github.com/gorilla/mux"
	"github.com/laughingcabbage/golinks/blockchain"
)

// Info contains the version, status of this program
type Info struct {
	Version string `json:"version"`
}

const (
	chainName = "vicechain"
)

var (
	info          Info
	chain         blockchain.Blockchain
	viceChainPath string

	govicePath = os.Getenv("GOVICE")
)

// See setup.go
func init() {
	checkEnv()
	loadInfo()
	loadBlockchain()
}

func main() {
	log.Println("Running...")

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheckHandler).Methods("GET")
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.HandleFunc("/chain", getChainHandler).Methods("GET")
	// router.HandleFunc("/chain", postBlockHandler).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode("UP"); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.Fatal(err)
	}
}

func getChainHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(chain); err != nil {
		log.Fatal(err)
	}
}

//TODO change to post chain. Post chain and compare with existing copy.
func postChainHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var blk block.Block
	if err := json.Unmarshal(b, &blk); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
