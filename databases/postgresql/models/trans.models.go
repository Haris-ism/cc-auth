package models

import "gorm.io/gorm"

type TableCreditCards struct{
	gorm.Model
	Bank			string			`gorm:"column:bank"`
	Limit			int				`gorm:"column:limit"`
	Balance			int				`gorm:"column:balance"`
	CC_Number		string			`gorm:"column:cc_number"`
	CVV				string			`gorm:"column:cvv"`
	CredsEmail		string			`gorm:"column:creds_email"`
	Credentials		Credentials		`gorm:"foreignKey:creds_email;references:email" json:"credentials,omitempty"`
}

func (t *TableCreditCards) TableName()string{
	return "credit_cards"
}
type CreditCards struct{
	// gorm.Model
	ID				int				`gorm:"column:id"`
	Bank			string			`gorm:"column:bank"`
	Limit			int				`gorm:"column:limit"`
	Balance			int				`gorm:"column:balance"`
	CC_Number		string			`gorm:"column:cc_number"`
	CVV				string			`gorm:"column:cvv"`
	CredsEmail		string			`gorm:"column:creds_email"`
}

func (t *CreditCards) TableName()string{
	return "credit_cards"
}