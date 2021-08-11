package database

import (
	"time"
)

type Wallet struct {
	id    		string
	userId		string
	userName 	string
	asset_tag	string
	asset_name	string
	hash 		string  
	balance 	float64 
	createdAt 	time.Time
	updatedAt 	time.Time
}
