package server

import (
	"context"

	"github.com/matheusrocha-mb/goproject/internal/database"
	wlt "github.com/matheusrocha-mb/goproject/pkg"
	log "github.com/sirupsen/logrus"
)

type ManagerWalletOp func(s *Server, cast database.Wallet) (*wlt.ModelWalletReply, error)

var walletOption = map[string]ManagerWalletOp{
	"CREATE": WalletCreate,
	"READ":   WalletRead,
	"DELETE": WalletDelete,
	"UPDATE": WalletUpdate,
}

func (s *Server) ManagerWallet(ctx context.Context, input *wlt.WalletRequest) (*wlt.ModelWalletReply, error) {
	log.Info("Start Manage function...")

	results := wlt.ModelWalletReply{}

	return &results, nil
}

func WalletCreate(s *Server, cast database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply
	//	var d database.Wallet

	return result, nil
}

func WalletRead(s *Server, cast database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply

	return result, nil
}

func WalletUpdate(s *Server, cast database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply

	return result, nil
}

func WalletDelete(s *Server, cast database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply

	return result, nil
}

/*
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	var wallet database.Wallet
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
*/
