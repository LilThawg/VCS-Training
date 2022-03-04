package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header["Token"] == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No Token Found",
			})
			c.Abort()
			return
		}

		var mySigningKey = []byte("secretkey")

		token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing token")
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Your Token has been expired.",
			})
			c.Abort()
			return
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["role"])
			if claims["role"] == "admin" {
				c.Next()
			} else {
				c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
					"error": "User not allowed",
				})
				c.Abort()
			}
		} 

	}
}