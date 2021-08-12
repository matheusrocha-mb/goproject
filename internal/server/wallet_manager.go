package server

import (
	"context"
	"fmt"

	"github.com/matheusrocha-mb/goproject/internal/database"
	wlt "github.com/matheusrocha-mb/goproject/pkg"
	log "github.com/sirupsen/logrus"
)

type ManagerWalletOp func(s *Server, obj database.Wallet) (*wlt.ModelWalletReply, error)

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

func WalletCreate(s *Server, obj database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply
	var d database.Wallet

	return result, nil
}

func WalletRead(s *Server, obj database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply
	var d database.Wallet

	d, err := s.db.GetWallet(obj.userId)
	if err != nil {
		return result, err
	}
	if (d != database.Wallet{}) {
		result.Data = &wlt.ModelWallet{
			UserId:    d.userId,
			UserName:  d.userName,
			AssetTag:  d.asset_tag,
			AssetName: d.asset_name,
			Hash:      d.Hash,
			Balance:   d.balance,
			CreatedAt: d.createdAt.string(),
			UpdatedAt: d.updatedAt.string(),
		}
	} else {
		result.Data = &wlt.ModelWallet{
			UserId: d.userId,
		}
		err = fmt.Errorf("User does not exists")
	}
	return result, err
}

func WalletUpdate(s *Server, obj database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply
	var d database.Wallet

	d, err := s.db.GetWallet(obj.userId)
	if err != nil {
		return result, err
	}
	if (d != database.Wallet{}){
		
	}
	return result, nil
}

func WalletDelete(s *Server, obj database.Wallet) (*wlt.ModelWalletReply, error) {
	var result *wlt.ModelWalletReply

	return result, nil
}