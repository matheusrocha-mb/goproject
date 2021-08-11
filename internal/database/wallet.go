package database

import (
	"time"
)

func (d *database_struct) InsertWallet(userId string, userName string, asset_tag string, asset_name string, hash string, balance float64) (Wallet, error) {
	var createdAt, updatedAt = time.Now(), time.Now()

	query := "INSERT INTO wallets (userId, userName, asset_tag, asset_name, hash, balance, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

	_, err := d.conn.Exec(query, userId, userName, asset_tag, asset_name, hash, balance, createdAt, updatedAt)
	if err != nil {
		d.log.Errorf("Database insert error: %v", err)
		return Wallet{}, err
	}
	return Wallet{
		userId:     userId,
		userName:   userName,
		asset_tag:  asset_tag,
		asset_name: asset_name,
		hash:       hash,
		balance:    balance,
		createdAt:  createdAt,
		updatedAt:  updatedAt}, nil
}

func (d *database_struct) GetWallet(userId string) (Wallet, error) {
	query := "SELECT userId, userName, asset_tag, asset_name, hash, balance, createdAt, updatedAt FROM wallets WHERE userId = $1;"

	current_wallet := Wallet{}

	if err := d.conn.QueryRow(query, userId).Scan(&current_wallet.userId, &current_wallet.userName,
		&current_wallet.asset_tag, &current_wallet.asset_name, &current_wallet.hash, &current_wallet.balance,
		&current_wallet.createdAt, &current_wallet.updatedAt); err != nil {
		d.log.Errorf("Database read error: %v", err)
		return current_wallet, nil
	}

	return current_wallet, nil
}

func (d *database_struct) DeleteWallet(userId string) (int64, error) {
	query := "DELETE FROM wallets WHERE userId = $1;"

	rm, err := d.conn.Exec(query, userId)
	if err != nil {
		d.log.Errorf("Database delete error: %v", err)
		return -1, err
	}

	return rm.RowsAffected()
}

func (d *database_struct) UpdateWalletUserName(userId string, userName string) (int64, error) {
	query := "UPDATE wallets SET userName = $1, updatedAt = $2 WHERE userId = $3;"
	t := time.Now()

	queryset, err := d.conn.Exec(query, userName, t, userId)
	if err != nil {
		d.log.Errorf("%v", err)
		return -1, nil
	}
	return queryset.RowsAffected()
}

func (d *database_struct) UpdateWalletAsset(userId string, asset_tag string, asset_name string) (int64, error) {
	query := "UPDATE wallets SET asset_tag = $1, asset_name = $2, updatedAt = $3 WHERE userId = $4;"
	t := time.Now()

	queryset, err := d.conn.Exec(query, asset_tag, asset_name, t, userId)
	if err != nil {
		d.log.Errorf("%v", err)
		return -1, nil
	}
	return queryset.RowsAffected()
}

func (d *database_struct) UpdateWalletHash(userId string, hash string) (int64, error) {
	query := "UPDATE wallets SET hash = $1, updatedAt = $2 WHERE userId = $3;"
	t := time.Now()

	queryset, err := d.conn.Exec(query, hash, t, userId)
	if err != nil {
		d.log.Errorf("%v", err)
		return -1, nil
	}
	return queryset.RowsAffected()
}

func (d *database_struct) UpdateWalletBalance(userId string, balance float64) (int64, error) {
	query := "UPDATE wallets SET balance = $1, updatedAt = $2 WHERE userId = $3;"
	t := time.Now()

	queryset, err := d.conn.Exec(query, balance, t, userId)
	if err != nil {
		d.log.Errorf("%v", err)
		return -1, nil
	}
	return queryset.RowsAffected()
}
