package controllers

import (
	"github.com/challenge2/database"
	"github.com/challenge2/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 	Sign Up
// @Description Đăng ký
// @Tags 		Handler Auth
// @Accept		json
// @Produce     json
// @Param 		user body models.UserSignUp true "User Sign Up" 
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleFailure
// @Router      /SignUp [post]
func SignUp(c *gin.Context) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := validateUser(user)
	
	if message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	var dbuser models.User
	connection.Where("email = ?", user.Email).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already in use",
		})
		return
	}

	user.Password, _ = GeneratehashPassword(user.Password)

	//set user's default role is : user
	user.Role = "user"

	//insert user details in database
	connection.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign up successfully!",
	})
}

// @Summary 	Sign In
// @Description Đăng Nhập
// @Tags 		Handler Auth
// @Accept		json
// @Produce     json
// @Param 		user body models.UserSignIn true "User Sign In" 
// @Success		200 {object} models.Token
// @Failure     400 {object} models.ExampleFailure
// @Router      /SignIn [post]
func SignIn(c *gin.Context) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var authDetails models.Authentication
	if err := c.ShouldBindJSON(&authDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var authUser models.User
	
	connection.Where("email = ?", authDetails.Email).First(&authUser)

	if authUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or Password is incorrect",
		})
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is incorrect",
		})
		return
	}

	validToken, err := GenerateJWT(authUser.Email, authUser.Role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	var token models.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// @Summary 	Sign Out
// @Description Đăng Xuất
// @Tags 		Handler Auth
// @Success		200 {object} models.ExampleSuccess
// @Router      /SignOut [post]
func SignOut(c *gin.Context) {

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign out successfully!", 
	})
}