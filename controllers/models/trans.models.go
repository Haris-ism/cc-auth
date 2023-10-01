package models

import "gorm.io/gorm"

type ReqCC struct{
	gorm.Model
	Bank			string			`json:"bank"`
	Limit			int				`json:"limit"`
	Balance			int				`json:"balance"`
	Email			string			`json:"email"`
}