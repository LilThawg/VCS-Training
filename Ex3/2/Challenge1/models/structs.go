package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Name	string `json:"name" validate:"required"`
	Dob		string `json:"dob" validate:"required,checkdate"`
	Gender	string `json:"gender" validate:"required,oneof=male female"`
	Email	string `gorm:"unique" json:"email" validate:"required,email"`
}

type UserCreateParam struct{
	Name	string `json:"name"`
	Dob		string `json:"dob" `
	Gender	string `json:"gender"`
	Email	string `gorm:"unique"`
}

type ExampleError struct {
	Error string `json:"error" example:"Lỗi gì đó ..."`
}

type ExampleSuccess struct {
	Success string `json:"message" example:"Thông báo đã thực hiện thành công..."`
}