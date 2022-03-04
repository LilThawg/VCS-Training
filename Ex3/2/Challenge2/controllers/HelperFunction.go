package controllers

import (
	"regexp"
	"time"

	"github.com/challenge2/models"
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

	re := regexp.MustCompile(`^[a-zA-Z0-9\-\_]*$`)

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
