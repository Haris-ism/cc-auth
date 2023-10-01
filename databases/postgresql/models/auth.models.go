package models

import "gorm.io/gorm"

type Credentials struct{
	gorm.Model
	Email		string	`gorm:"unique"`
	Password	string	`gorm:"column:password"`
}

func (t *Credentials) TableName()string{
	return "credentials"
}

