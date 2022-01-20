package controllers

import (
	"auth-jwt/database"
	"auth-jwt/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	
	var users []models.User

	if result := connection.Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}

	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	// Read the email parameter
	var user models.User
	// fmt.Println(user)
	err := json.NewDecoder(r.Body).Decode(&user)
	// fmt.Println(user)
	if err != nil {
		var err models.Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	email := user.Email
	fmt.Println(email)
	// connection.Unscoped().Where("email = ?", email).Delete(&models.User{})
	// fmt.Println(user)
	// fmt.Println(models.User{})
	connection.Unscoped().Where("email = ?", email).Delete(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}

	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err models.Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}
	
	connection.Create(&user)

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}