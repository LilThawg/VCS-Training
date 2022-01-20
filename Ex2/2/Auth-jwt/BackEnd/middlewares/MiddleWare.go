package middlewares

import (
	"auth-jwt/controllers"
	"auth-jwt/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

//---------------------MIDDLEWARE FUNCTION-----------------------

//check whether user is authorized or not
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			var err models.Error
			err = controllers.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte("secretkey")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing token")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err models.Error
			err = controllers.SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return

			}
		}
		var reserr models.Error
		reserr = controllers.SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}