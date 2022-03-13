package controllers

import (
	"github.com/challenge2/database"
	"github.com/challenge2/models"

	"net/http"

	"github.com/gin-gonic/gin"
)


// AddUser godoc
// @Summary 	Add a user
// @Description	Thêm 1 User
// @Tags		Handler User (Only Admin)
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Param 		user body models.UserAdd true "Add User"
// @Success		200 {object} models.User
// @Failure     400 {object} models.ExampleFailure
// @Router      /admin/addUser [post]
func AddUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	
	message := validateUser(newUser)
	
	if message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	
	var dbuser models.User
	connection.Where("email = ?", newUser.Email).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already in use",
		})
		return
	}

	newUser.Password, _ = GeneratehashPassword(newUser.Password)

	c.JSON(http.StatusOK, newUser)
	connection.Create(&newUser)
	
}

// GetAllUsers godoc
// @Summary 	Get All Users
// @Description	Lấy thông tin toàn bộ user
// @Tags		Handler User (Only Admin)
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Success		200 {array} models.UserAdd
// @Failure     400 {object} models.ExampleFailure
// @Router      /admin/getAllUsers [get]
func GetAllUsers(c *gin.Context) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	
	var users []models.User

	if result := connection.Find(&users); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	} 
	
	c.JSON(http.StatusOK, users)
	
}

// DeleteUser godoc
// @Summary 	Delete A User
// @Description	Xoá thông tin 1 user
// @Tags		Handler User (Only Admin)
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Param 		email query string true "User's Email" 
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleFailure
// @Router      /admin/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	
	var user models.User
	email := c.Query("email") 
	
	user.Email = email
	connection.Unscoped().Where("email = ?", email).Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted!",
	})
}

// UpdateUser godoc
// @Summary 	Update A User
// @Description	Chỉnh thông tin 1 user
// @Tags		Handler User (Only Admin)
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Param 		email query string true "User's Email" 
// @Param 		user body models.UserAdd true "Update User"
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleFailure
// @Router      /admin/updateUser [put]
func UpdateUser(c *gin.Context) {
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	message := validateUser(updatedUser)
	
	if message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var user models.User
	email := c.Query("email")
	updatedUser.Password, _ = GeneratehashPassword(updatedUser.Password)

	connection.Model(&user).Where("email = ?", email).Update(&updatedUser)
	


	c.JSON(http.StatusOK, gin.H{
		"message": "Updated!",
	})
	
}
