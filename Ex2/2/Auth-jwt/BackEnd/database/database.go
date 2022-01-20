package database

import (
	"auth-jwt/models"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//-------------DATABASE FUNCTIONS---------------------

//returns database connection
func GetDatabase() *gorm.DB {
	databasename := "authServer"
	database := "postgres"
	databasepassword := "lethang2909"
	databaseurl := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"

	connection, err := gorm.Open(database, databaseurl)
	if err != nil {
		log.Fatalln("Invalid database url")
	}
	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("Database connected")
	}
	fmt.Println("Database connection successful.")
	return connection
}

//create user table in userdb
func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(models.User{})
}

//closes database connection
func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
