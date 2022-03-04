package main

import (
	"github.com/challenge1/database"
	"github.com/challenge1/routes"
)

// Khai báo Swagger

// @title           Swagger Example API
// @version         1.0
// @description     Đây là swagger api của challenge 1
// @host localhost:8080
// @BasePath /api

func main() {
	database.InitialMigration()
	routes.ServerStart()
}