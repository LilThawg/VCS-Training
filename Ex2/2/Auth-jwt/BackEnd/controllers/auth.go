package controllers

import (
	"auth-jwt/models"
	"auth-jwt/database"
	"encoding/json"
	"log"
	"net/http"
	// "fmt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
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

	var dbuser models.User
	connection.Where("email = ?", user.Email).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Email != "" {
		var err models.Error
		err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	//set user's default role is : user
	user.Role = "user"
	
	//insert user details in database
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var authDetails models.Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err models.Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser models.User
	
	connection.Where("email = ?", authDetails.Email).First(&authUser)
	
	if authUser.Email == "" {
		var err models.Error
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err models.Error
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(authUser.Email, authUser.Role)
	if err != nil {
		var err models.Error
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token models.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w , r, "/", http.StatusFound)
}