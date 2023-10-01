package usecase

import (
	"cc-auth/controllers/models"
	trans "cc-auth/databases/postgresql/models"
	"cc-auth/utils"
	"errors"
)


func (uc *usecase)AddCC(req models.ReqCC)error{
	cc:=trans.CreditCards{}
	cred,err:=uc.postgre.QueryEmailCC(req.Email)
	if err!=nil{
		return err
	}

	if cred.ID!=0{
		return errors.New("Email Already Registered")
	}

	cc_num,err:=utils.GenerateRandom(10)
	if err!=nil{
		return errors.New("Failed to Generate CC")
	}

	cvv,err:=utils.GenerateRandom(3)
	if err!=nil{
		return errors.New("Failed to Generate CVV")
	}

	cc.Bank=req.Bank
	cc.CredsEmail=req.Email
	cc.Limit=req.Limit
	cc.CC_Number=cc_num
	cc.CVV=cvv
	
	err=uc.postgre.AddCC(cc)

	if err!=nil{
		return err
	}

	return nil
}

func (uc *usecase)TopUpCC(req models.ReqCC)error{
	cred,err:=uc.postgre.QueryEmailCC(req.Email)
	if err!=nil{
		return err
	}

	if cred.ID==0{
		return errors.New("Invalid Email")
	}

	cred.Balance+=req.Balance

	err=uc.postgre.TopUpCC(cred)
	if err!=nil{
		return err
	}

	return nil
}