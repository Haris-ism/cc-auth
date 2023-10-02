package postgre

import (
	"cc-auth/databases/postgresql/models"
)



func (db *postgreDB)TopUpCC(cred models.CreditCards)error{
	err:=db.postgre.Save(&cred)
	if err!=nil{
		return err.Error
	}
	return nil
}

