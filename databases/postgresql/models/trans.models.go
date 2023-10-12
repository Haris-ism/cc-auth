package models

import "gorm.io/gorm"

type CreditCards struct{
	gorm.Model
	Bank			string			`gorm:"column:bank"`
	Limit			int				`gorm:"column:limit"`
	Balance			int				`gorm:"column:balance"`
	CC_Number		string			`gorm:"column:cc_number"`
	CVV				string			`gorm:"column:cvv"`
	CredsEmail		string			`gorm:"column:creds_email"`
	Credentials		Credentials		`gorm:"foreignKey:creds_email;references:email" json:"credentials,omitempty"`
	// Credentials		Credentials		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (t *CreditCards) TableName()string{
	return "credit_cards"
}