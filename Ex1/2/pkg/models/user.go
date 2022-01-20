package models

type User struct{
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Dob	string `json:"dob"`
	Gender string `json:"gender"`
	Email string `json:"email"`
}
