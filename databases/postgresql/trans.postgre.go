package postgre

import (
	"cc-auth/databases/postgresql/models"
	"fmt"
)

func (db *postgreDB)AddCC(cc models.CreditCards)error{
	err:=db.postgre.Create(&cc).Error
	if err!=nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func (db *postgreDB)QueryEmailCC(email string)(models.CreditCards,error){
	cc:=models.CreditCards{}
	err:=db.postgre.Where("creds_email = ?",email).Find(&cc).Error
	if err!=nil{
		return cc,err
	}
	return cc,nil
}

func (db *postgreDB)TopUpCC(cred models.CreditCards)error{
	err:=db.postgre.Save(&cred)
	if err!=nil{
		return err.Error
	}
	return nil
}