package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/matheusrocha-mb/goproject/internal/handlers"
)
func Run(hport string) {

}

func httpServer(port string) *http.Server {
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

}