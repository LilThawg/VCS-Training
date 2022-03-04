package controllers

import (
	"net/http"

	"github.com/challenge1/database"
	"github.com/challenge1/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func checkDate(f1 validator.FieldLevel) bool {
	str := f1.Field().String()

	re := regexp.MustCompile(`^([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}$`)

	return re.MatchString(str)
}

// AddUser godoc
// @Summary 	Add a user
// @Description	Thêm 1 User
// @Security 	ApiKeyAuth
// @Tags		Handler User
// @Accept		json
// @Produce 	json	
// @Param 		user body models.UserCreateParam true "Add User"
// @Success		200 {object} models.User
// @Failure     400 {object} models.ExampleError
// @Router      /addUser [post]

func AddUser(c *gin.Context) {
	var validate *validator.Validate = validator.New()
	validate.RegisterValidation("checkdate", checkDate)
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var message string
	if err := validate.Struct(newUser); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				message = message + "Nhập thiếu thông tin. "
			}
			if err.ActualTag() == "checkdate" {
				message = message + "Nhập sai định dạng ngày tháng (dd/mm/yyyy). "
			}
			if err.ActualTag() == "oneof" {
				message = message + "Nhập sai giới tính (male/female). "
			}
			if err.ActualTag() == "email" {
				message = message + "Nhập sai định dạng email. "
			}
		}
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
	c.JSON(http.StatusOK, newUser)
	connection.Create(&newUser)
	
}

// GetAllUsers godoc
// @Summary 	Get All Users
// @Description	Lấy thông tin toàn bộ user
// @Tags		Handler User
// @Accept		json
// @Produce 	json	
// @Success		200 {array} models.UserCreateParam
// @Failure     400 {object} models.ExampleError
// @Router      /getAllUsers [get]
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
// @Tags		Handler User
// @Accept		json
// @Produce 	json	
// @Param 		email query string true "User's Email" 
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleError
// @Router      /deleteUser [delete]
func DeleteUser(c *gin.Context) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	
	var user models.User
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// } 
	
	email := c.Query("email")
	connection.Model(&user).Where("email = ?", email).Delete(&user)
	// connection.Unscoped().Where("email = ?", email).Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted!",
	})
}

// UpdateUser godoc
// @Summary 	Update A User
// @Description	Chỉnh thông tin 1 user
// @Tags		Handler User
// @Accept		json
// @Produce 	json	
// @Param 		email query string true "User's Email" 
// @Param 		user body models.UserCreateParam true "Add User"
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleError
// @Router      /updateUser [put]
func UpdateUser(c *gin.Context) {
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var user models.User
	email := c.Query("email")
	
	connection.Model(&user).Where("email = ?", email).Update(&updatedUser)

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated!",
	})
	
}