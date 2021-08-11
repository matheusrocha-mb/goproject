package database

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type DatabaseRepo interface {
	// Wallets
	InsertWallet(wallet Wallet) error
	GetWallet(wallet Wallet) (Wallet, error)
	UpdateWallet(wallet Wallet) (Wallet, error)
	DeleteWallet(wallet Wallet) (Wallet, error)

}

type database_struct struct {
	conn *sql.DB
	log  *logrus.Entry
}
