package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Wallet struct {
	ID    string `json:"id"`
	Asset *Asset `json:"asset"`
}
type Asset struct {
	Hash    string  `json:"hashid"`
	Balance float64 `json:"balance"`
}

var wallets []Wallet

const SERVER_LOCAL_PORT = 5002
const SERVER_CONTAINER_PORT = 5001
const RANDOM_NUMBER_RANGE = 999999999999999999

func main() {
	fmt.Println("Starting Go project...")

	wallets = append(wallets, Wallet{ID: "1", Asset: &Asset{Hash: "4e8sES76EAS", Balance: 3.14}})
	wallets = append(wallets, Wallet{ID: "2", Asset: &Asset{Hash: "4e8sES79PKE", Balance: 200.14}})
	r := mux.NewRouter()

	r.HandleFunc("/wallet", ListWallets).Methods("GET")
	r.HandleFunc("/wallet/{id}", GetWallet).Methods("GET")
	r.HandleFunc("/wallet/{id}/assets", ListAssets).Methods("GET")
	r.HandleFunc("/wallet/{id}/asset/{id}", GetAsset).Methods("GET")
	r.HandleFunc("/wallet/create", CreateWallet).Methods("POST")
	r.HandleFunc("/wallet/{id}/create", CreateAsset).Methods("POST")
	r.HandleFunc("/wallet/delete", DeleteWallet).Methods("DELETE")
	r.HandleFunc("/wallet/{id}/delete", DeleteAsset).Methods("DELETE")
	r.HandleFunc("/wallet/update", UpdateWallet).Methods("PUT")
	r.HandleFunc("/wallet/{id}/update", UpdateAsset).Methods("PUT")

	fmt.Printf("\nStarting server local on port: %d\nStarting server on container at port: %d\n", SERVER_LOCAL_PORT, SERVER_CONTAINER_PORT)
	log.Fatal(http.ListenAndServe(":5002", r))

}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	var wallet Wallet
	_ = json.NewDecoder(r.Body).Decode(&wallet)
	wallet.ID = strconv.Itoa(rand.Intn(RANDOM_NUMBER_RANGE))
	wallets = append(wallets, wallet)
	json.NewEncoder(w).Encode(wallet)
}

func ListWallets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(wallets)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)

	for _, wallet := range wallets {
		if wallet.ID == params["id"] {
			json.NewEncoder(w).Encode(wallet)
			return
		}
	}
}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)

	for _, wallet := range wallets {
		if wallet.ID == params["id"] {
			updatebalance, err := strconv.ParseFloat(params["balance"], 32)
			if err != nil {
				fmt.Println("Payload com balance Invalido.")
				return
			}
			wallet.Asset.Balance = updatebalance
			return
		}
	}
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)

	for index, wallet := range wallets {
		if wallet.ID == params["id"] {
			wallets = append(wallets[:index], wallets[index+1:]...)
		}
	}
}

func CreateAsset(w http.ResponseWriter, r *http.Request) {

}

func ListAssets(w http.ResponseWriter, r *http.Request) {

}

func GetAsset(w http.ResponseWriter, r *http.Request) {

}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {

}

func DeleteAsset(w http.ResponseWriter, r *http.Request) {

}
