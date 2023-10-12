package models

import "gorm.io/gorm"

type ReqCC struct{
	gorm.Model
	Bank			string			`json:"bank"`
	Limit			int				`json:"limit"`
	Balance			int				`json:"balance"`
	Email			string			`json:"email"`
}

type ReqItems struct {
	gorm.Model
	CC_Number		string			`json:"cc_number"`
	CVV				string			`json:"cvv"`
	Item			string			`json:"item"`
	Discount		string			`json:"discount"`
	Quantity		string			`json:"quantity"`
}