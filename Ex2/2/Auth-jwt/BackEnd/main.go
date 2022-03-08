package main

import (
	"auth-jwt/routes"
	"auth-jwt/database"
)

func main() {
	database.InitialMigration()
	routes.ServerStart()
}
