package usecase

import (
	"cc-auth/constants"
	"cc-auth/controllers/models"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) Register(req models.Credentials) error{
	cred,err:=uc.postgre.EmailQuery(req.Email)

	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}

	if cred.ID!=0{
		return errors.New(constants.EMAIL_REGISTERED)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		log.Println("failed to hash password")
	}

	cred.Email=req.Email
	cred.Password=string(hash)

	err=uc.postgre.CreateCredentials(cred)

	if err!=nil{
		return err
	}

	return nil
}

func (uc *usecase)Login (req models.Credentials)(string,error){
	cred,err:=uc.postgre.EmailQuery(req.Email)
	
	if err!=nil{
		return "",errors.New(constants.ERROR_DB)
	}

	if cred.ID==0{
		return "",errors.New(constants.INVALID_EMAIL)
	}

	err=bcrypt.CompareHashAndPassword([]byte(cred.Password),[]byte(req.Password))

	if err!=nil{
		return "",errors.New(constants.INVALID_INPUT)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))

	if err!=nil{
		return tokenString,errors.New("Failed to Generate Token")
	}

	return tokenString,nil
}

func (uc *usecase)DelCC(id string)error{
	ids,_:=strconv.Atoi(id)
	cred,err:=uc.postgre.QueryIDCC(ids)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}
	if cred.ID==0{
		return errors.New("Invalid Credit Card")
	}
	err=uc.postgre.DelCC(ids)
	if err!=nil{
		return errors.New("Failed to Delete Credit Card")
	}
	return nil
}