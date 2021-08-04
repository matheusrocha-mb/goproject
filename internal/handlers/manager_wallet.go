package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/matheusrocha-mb/goproject/internal/database"
	"github.com/matheusrocha-mb/goproject/main.go"
	"github.com/gorilla/mux"
)



const SERVER_LOCAL_PORT = 5002
const SERVER_CONTAINER_PORT = 5001
const RANDOM_NUMBER_RANGE = 999999999999999999


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