package controllers

import (
	"regexp"
	"strconv"
	"time"

	"github.com/challenge3/database"
	"github.com/challenge3/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretkey string = "secretkey"

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
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		// fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}


func checkString(f1 validator.FieldLevel) bool {
	str := f1.Field().String()

	re := regexp.MustCompile(`^[a-zA-Z0-9\-\_\s]+$`)

	return re.MatchString(str)
}

func validateUser(newUser models.User) string {
	var validate *validator.Validate = validator.New()
	validate.RegisterValidation("checkstring", checkString)
	var message string
	if err := validate.Struct(newUser); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				message = message + "Nhập thiếu thông tin. "
			}
			if err.ActualTag() == "checkstring" {
				message = message + "Tài khoản hoặc mật khẩu chỉ được dùng [a-Z], -, _, [0-9]. "
			}
			if err.ActualTag() == "email" {
				message = message + "Nhập sai định dạng email. "
			}
		}
	}
	return message
}

func validatePost(newPost models.Post) string {
	var validate *validator.Validate = validator.New()
	validate.RegisterValidation("checkstring", checkString)
	var message string
	if err := validate.Struct(newPost); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				message = message + "Nhập thiếu thông tin. "
			}
			if err.ActualTag() == "checkstring" {
				message = message + "Title hoặc content chỉ được dùng [a-Z], -, _, [0-9]. "
			}
		}
	}
	return message
}

//GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	// Initializing default
	//	var mode string
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
			case "limit":
				limit, _ = strconv.Atoi(queryValue)
				break
			case "page":
				page, _ = strconv.Atoi(queryValue)
				break
			case "sort":
				sort = queryValue
				break
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}

func GetPaginationPost(post *models.Post, pagination *models.Pagination) (*[]models.Post,int64,error) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var posts []models.Post
	var totalRows int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Post{}).Where(post).Find(&posts)
	if result.Error != nil {
		msg := result.Error
		return nil,totalRows, msg
	}
	errCount := connection.Model(&models.Post{}).Count(&totalRows).Error
	if errCount != nil {
		return nil, totalRows, errCount
	}
	return &posts, totalRows, nil
}