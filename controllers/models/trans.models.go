package models

import "gorm.io/gorm"

type ReqCC struct{
	gorm.Model
	Bank			string			`json:"bank"`
	Limit			int				`json:"limit"`
	Balance			int				`json:"balance"`
	Email			string			`json:"email"`
}

type TransactionItems struct{
	ItemID			int			`json:"item_id"`
	Discount		string		`json:"discount"`
	Quantity		int			`json:"quantity"`
	CCNumber		string		`json:"cc_number"`
	CVV				string		`json:"cvv"`
	Amount			int			`json:"amount"`
	Price			int			`json:"price"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Percentage		int			`json:"percentage"`
}

