package usecase

import (
	"cc-auth/constants"
	"cc-auth/controllers/models"
	trans "cc-auth/databases/postgresql/models"
	tModels "cc-auth/hosts/transaction/models"
	"cc-auth/middleware"
	"context"
	"encoding/json"
	"log"
	"time"

	"cc-auth/utils"
	"errors"

	"google.golang.org/grpc/metadata"
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
	result:=tModels.DecTransItem{}
	encryptedReq,err:=utils.EncryptTransItemGrpc(req)
	if err!=nil{
		log.Println("error encrypt:",err)
		return result, err
	}

	bytes,err:=json.Marshal(encryptedReq)
	if err!=nil{
		log.Println("error marshal:",bytes)
		return result,err
	}

	signature:=middleware.Signature(string(bytes),timeStamp)
	meta:=map[string]string{
		"timestamp":timeStamp,
		"signature":signature,
	}
	md := metadata.New(meta)
	ctx:=metadata.NewOutgoingContext(context.Background(),md)
	res,err:=uc.hostGrpc.Transaction().TransItems(ctx,encryptedReq)
	if err!=nil{
		log.Println("err grpc req:",res)
		return result, err
	}

	resp,err:=utils.DecryptTransItemRes(res.Data)
	if err!=nil{
		log.Println("err decrypt:",err)
		return result, err
	}

	return resp,nil
}
