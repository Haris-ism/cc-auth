package postgre

import (
	"cc-auth/controllers/models"
	dbs "cc-auth/databases/postgresql/models"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *postgreDB) Register(req models.Credentials)error{
	cred:=dbs.Credentials{}
	if err:=db.postgre.Where("email = ?",req.Email).Find(&cred).Error;err!=nil{
		return err
	}

	if cred.ID!=0{
		return errors.New("email already exist")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Println("failed to hash password")
	}

	cred.Email=req.Email
	cred.Password=string(hash)

	if err:=db.postgre.Create(&cred).Error;err!=nil{
		return err
	}

	return nil
}