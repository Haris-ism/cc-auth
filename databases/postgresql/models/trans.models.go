package models

import "gorm.io/gorm"

type CreditCards struct{
	gorm.Model
	Bank			string			`gorm:"column:bank"`
	Limit			int				`gorm:"column:limit"`
	Balance			int				`gorm:"column:balance"`
	CC_Number		int				`gorm:"column:cc_number"`
	CVV				int				`gorm:"column:cvv"`
	CredsEmail			string		`gorm:"column:creds_email"`
	Credentials		Credentials		`gorm:"foreignKey:creds_email;references:email"`
	// Credentials		Credentials		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (t *CreditCards) TableName()string{
	return "credit_cards"
}