package models

import "gorm.io/gorm"

type TableCredentials struct{
	gorm.Model
	Email		string	`json:"email,omitempty" gorm:"unique"`
	Password	string	`json:"password,omitempty" gorm:"column:password"`
}

func (t *TableCredentials) TableName()string{
	return "credentials"
}

type Credentials struct{
	gorm.Model
	Email		string	`json:"email,omitempty" gorm:"unique"`
	Password	string	`json:"password,omitempty" gorm:"column:password"`
}

func (t *Credentials) TableName()string{
	return "credentials"
}

