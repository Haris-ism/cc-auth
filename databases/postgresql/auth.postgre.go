package postgre

import (
	"cc-auth/controllers/models"
	auth "cc-auth/controllers/models"
	dbs "cc-auth/databases/postgresql/models"
)

func (db *postgreDB) EmailQuery(email string)(dbs.Credentials,error){
	cred:=dbs.Credentials{}
	err:=db.postgre.Where("email = ?",email).Find(&cred).Error

	if err!=nil{
		return cred,err
	}

	return cred,nil
}

func (db *postgreDB) CreateCredentials(cred dbs.Credentials)error{
	err:=db.postgre.Create(&cred).Error

	if err!=nil{
		return err
	}

	return nil
}

func (dbs *postgreDB) Login(req models.Credentials) error{
	cred:=auth.Credentials{}

	if err:=dbs.postgre.Where("email = ?",req.Email).Find(&cred).Error; err!=nil{
		return err
	}

	return nil
}

func (db *postgreDB)AddCC(cc dbs.TableCreditCards)error{
	err:=db.postgre.Create(&cc).Error
	if err!=nil{
		return err
	}
	return nil
}

func (db *postgreDB)QueryEmailCC(email string)(dbs.CreditCards,error){
	cc:=dbs.CreditCards{}
	err:=db.postgre.Where("creds_email = ?",email).Find(&cc).Error
	if err!=nil{
		return cc,err
	}
	return cc,nil
}

func (db *postgreDB)QueryIDCC(id int)(dbs.CreditCards,error){
	cred:=dbs.CreditCards{}
	err:=db.postgre.Where("ID = ?", id).Find(&cred).Error
	if err!=nil{
		return cred,err
	}
	return cred,nil
}

func (db *postgreDB)DelCC(id int)error{
	cred:=dbs.CreditCards{}
	err:=db.postgre.Where("ID = ?",id).Delete(&cred).Error
	if err!=nil{
		return err
	}
	return nil
}

func (db *postgreDB)GetCC()([]dbs.CreditCards,error){
	cred:=[]dbs.CreditCards{}
	err:=db.postgre.Find(&cred).Error
	if err!=nil{
		return cred,err
	}
	return cred,nil
}