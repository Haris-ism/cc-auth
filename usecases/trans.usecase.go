package usecase

import (
	"cc-auth/constants"
	"cc-auth/controllers/models"
	trans "cc-auth/databases/postgresql/models"
	tModels "cc-auth/hosts/transaction/models"
	"cc-auth/middleware"
	"encoding/json"
	"time"

	"cc-auth/utils"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)


func (uc *usecase)AddCC(req models.ReqCC)error{
	cc:=trans.TableCreditCards{}
	cred,err:=uc.postgre.QueryEmailCC(req.Email)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}

	if cred.ID!=0{
		return errors.New(constants.EMAIL_REGISTERED)
	}

	cc_num,err:=utils.GenerateRandom(10)
	if err!=nil{
		return errors.New(constants.ERROR_CC)
	}

	cvv,err:=utils.GenerateRandom(3)
	if err!=nil{
		return errors.New(constants.ERROR_CVV)
	}

	cc.Bank=req.Bank
	cc.CredsEmail=req.Email
	cc.Limit=req.Limit
	cc.CC_Number=cc_num
	cc.CVV=cvv
	
	err=uc.postgre.AddCC(cc)

	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}

	return nil
}

func (uc *usecase)TopUpCC(req models.ReqCC)error{
	cred,err:=uc.postgre.QueryEmailCC(req.Email)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}

	if cred.ID==0{
		return errors.New(constants.INVALID_EMAIL)
	}

	cred.Balance+=req.Balance

	err=uc.postgre.TopUpCC(cred)
	if err!=nil{
		return errors.New(constants.ERROR_DB)
	}

	return nil
}

func (uc *usecase)GetCC()([]trans.CreditCards,error){
	cred,err:=uc.postgre.GetCC()
	if err!=nil{
		return cred,errors.New(constants.ERROR_DB)
	}
	return cred,nil
}

func (uc *usecase)TransItem(req tModels.TransactionItems)(tModels.DecTransItem,error){
	timeStamp:=time.Now().Format("15:04:05")
	result:=tModels.ResponseTransactionItems{}
	req,err:=utils.EncryptTransItem(req)
	if err!=nil{
		return result.Data, err
	}
	bytes,err:=json.Marshal(req)
	if err!=nil{
		return result.Data,err
	}

	signature:=middleware.Signature(string(bytes),timeStamp)
	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	header.Add("TimeStamp", timeStamp)
	header.Add("Signature", signature)
	res,bytes,err:=uc.host.Transaction().Send(constants.TRANSACTION_ITEMS,req,header)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_REQUEST_FAILED)
	}
	
	resHost:=tModels.ResHostTransactionItems{}
	err=json.Unmarshal(bytes,&resHost)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_REQUEST_FAILED)
	}
	if res.StatusCode!=200{
		return result.Data, errors.New(resHost.Message)
	}
	resp,err:=utils.DecryptTransItemRes(resHost.Data)
	if err!=nil{
		logrus.Error(err)
		return result.Data, err
	}

	return resp,nil
}


