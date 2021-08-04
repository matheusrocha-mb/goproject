package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusrocha-mb/goproject/internal/handlers"
)

const SERVER_LOCAL_PORT = 5002
const SERVER_CONTAINER_PORT = 5001

func Run(hport string) {
	fmt.Println()
	httpsrv := httpServer()

	fmt.Printf("\nStarting server local on port: %d\nStarting server on container at port: %d\n", SERVER_LOCAL_PORT, SERVER_CONTAINER_PORT)
	log.Fatal(http.ListenAndServe(":5002", httpsrv))
}

func httpServer() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/wallet", handlers.ListWallets).Methods("GET")
	r.HandleFunc("/wallet/{id}", handlers.GetWallet).Methods("GET")
	r.HandleFunc("/wallet/{id}/assets", handlers.ListAssets).Methods("GET")
	r.HandleFunc("/wallet/{id}/asset/{id}", handlers.GetAsset).Methods("GET")
	r.HandleFunc("/wallet/create", handlers.CreateWallet).Methods("POST")
	r.HandleFunc("/wallet/{id}/create", handlers.CreateAsset).Methods("POST")
	r.HandleFunc("/wallet/delete", handlers.DeleteWallet).Methods("DELETE")
	r.HandleFunc("/wallet/{id}/delete", handlers.DeleteAsset).Methods("DELETE")
	r.HandleFunc("/wallet/update", handlers.UpdateWallet).Methods("PUT")
	r.HandleFunc("/wallet/{id}/update", handlers.UpdateAsset).Methods("PUT")

	return r
}
