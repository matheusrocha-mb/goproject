package main

import (
	"fmt"

	"github.com/matheusrocha-mb/goproject/internal/database"
	"github.com/matheusrocha-mb/goproject/internal/server"
)

const RANDOM_NUMBER_RANGE = 999999999999999999

func main() {
	fmt.Println("Starting Go project...")
	server.Run("5002")

	database.ConnectDB()

}
