package usecase

import (
	"cc-auth/constants"
	"cc-auth/hosts/callback/models"
	"encoding/json"
	"errors"
	"log"
)

func (uc *usecase)InquiryItems()([]models.InquiryItems,error){
	result:=[]models.InquiryItems{}
	res,err:=uc.hostGrpc.Callback().InquiryItems()
	if err!=nil{
		log.Println("err:",err)
		return result, errors.New(constants.ERROR_DB)
	}
	if res.Code!=200{
		log.Println("err res code:",res.Code)
		return result, errors.New(constants.ERROR_INQUIRY)
	}
	log.Println("grpc res:",res)
	bytes,err:=json.Marshal(res.Data)
	if err!=nil{
		return result,errors.New(constants.ERROR_DB)
	}
	err=json.Unmarshal(bytes,&result)
	return result,nil
}
func (uc *usecase)InquiryDiscounts()([]models.InquiryDiscounts,error){
	result:=[]models.InquiryDiscounts{}
	res,err:=uc.hostGrpc.Callback().InquiryDiscounts()
	if err!=nil{
		log.Println("err:",err)
		return result, errors.New(constants.ERROR_DB)
	}
	if res.Code!=200{
		log.Println("err res code:",res.Code)
		return result, errors.New(constants.ERROR_INQUIRY)
	}
	log.Println("grpc res:",res)
	bytes,err:=json.Marshal(res.Data)
	if err!=nil{
		return result,errors.New(constants.ERROR_DB)
	}
	err=json.Unmarshal(bytes,&result)
	return result,nil
}