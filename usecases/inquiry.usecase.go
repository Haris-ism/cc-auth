package usecase

import (
	"cc-auth/constants"
	"cc-auth/hosts/callback/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (uc *usecase)InquiryItems()([]models.InquiryItems,error){
	// result:=models.ResponseMerchantItems{}
	// header := make(http.Header)
	// header.Add("Accept", "*/*")
	// header.Add("Content-Type", "application/json")
	// _,data,err:=uc.host.Callback().Send(constants.INQUIRY_ITEMS,"",header)
	// if err!=nil{
	// 	return result.Data, errors.New(constants.ERROR_DB)
	// }
	// err=json.Unmarshal(data,&result)
	// if err!=nil{
	// 	return result.Data, errors.New(constants.ERROR_INQUIRY)
	// }
	// if result.Code!=200{
	// 	return result.Data, errors.New(constants.ERROR_INQUIRY)
	// }
	// return result.Data, nil

	result:=[]models.InquiryItems{}
	log.Println("masuk uc")
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
	result:=models.ResponseMerchantDiscounts{}
	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	res,data,err:=uc.host.Callback().Send(constants.INQUIRY_DISCOUNTS,"",header)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_DB)
	}
	err=json.Unmarshal(data,&result)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	if res.StatusCode!=200{
		return result.Data, errors.New(result.Message)
	}
	return result.Data, nil
}