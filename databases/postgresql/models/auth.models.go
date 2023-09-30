package models

import "gorm.io/gorm"

type Credentials struct{
	gorm.Model
	Email		string	`gorm:"email"`
	Password	string	`gorm:"password"`
}