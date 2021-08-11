package main

import (
	"fmt"

	"github.com/matheusrocha-mb/goproject/internal/database"
	"github.com/matheusrocha-mb/goproject/internal/server"
)


func main() {
	fmt.Println("Starting Go project...")
	
	database.ConnectDB()

	server.Run("5002")

}
