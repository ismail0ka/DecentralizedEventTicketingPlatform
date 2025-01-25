package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        Role   `json:"role"`
	Password    string `json:"password"`
}
