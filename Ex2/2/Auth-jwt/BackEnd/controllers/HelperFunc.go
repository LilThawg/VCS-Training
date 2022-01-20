package controllers

import (
	"auth-jwt/models"
	"net/http"
	// "fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

//--------------HELPER FUNCTIONS---------------------
var secretkey string = "secretkey"

//set error message in Error struct
func SetError(err models.Error, message string) models.Error {
	err.IsError = true
	err.Message = message
	return err
}

//take password as input and generate new hash password from it
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Generate JWT token
func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		// fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
}

// func AdminIndex(w http.ResponseWriter, r *http.Request) {
// 	if r.Header.Get("Role") != "admin" {
// 		w.Write([]byte("Not authorized."))
// 		return
// 	}
// 	w.Write([]byte("Welcome, Admin."))
// }

// func UserIndex(w http.ResponseWriter, r *http.Request) {
// 	if r.Header.Get("Role") != "user" {
// 		w.Write([]byte("Not Authorized."))
// 		return
// 	}
// 	w.Write([]byte("Welcome, User."))
// }