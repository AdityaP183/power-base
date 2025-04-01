package main

import (
	"github.com/AdityaP183/power_base/internal/config"
	"github.com/AdityaP183/power_base/internal/database"
)

func main() {
	// loading env variables
	config.LoadEnv()

	// database connection
	database.ConnectDB()
}
