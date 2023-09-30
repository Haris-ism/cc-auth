package postgre

import (
	"cc-auth/controllers/models"
	auth "cc-auth/controllers/models"
	dbs "cc-auth/databases/postgresql/models"
)

func (db *postgreDB) EmailQuery(req models.Credentials)(dbs.Credentials,error){
	cred:=dbs.Credentials{}
	err:=db.postgre.Where("email = ?",req.Email).Find(&cred).Error

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