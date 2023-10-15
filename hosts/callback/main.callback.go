package callback

import (
	"cc-auth/constants"
	"cc-auth/utils"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type(
	callback struct{
		callbackHost		string
		inquiryItems		string
		inquiryDiscounts	string
	}
	CallbackInterface interface{
		Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error)
	}
)

func (h *callback)Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error){
	var url string
	var method string
	switch types{
		case constants.INQUIRY_ITEMS:
			url=h.callbackHost+h.inquiryItems
			method=constants.HTTP_GET
		case constants.INQUIRY_DISCOUNTS:
			url=h.callbackHost+h.inquiryDiscounts
			method=constants.HTTP_GET
	}
	res,data,err:=utils.HTTPRequest(url,method,body,header)
	
	if err!=nil{
		return res,data,err
	}
	return res,data,nil
}

func InitCallback()CallbackInterface{
	return &callback{
		callbackHost:utils.GetEnv("CALLBACK_HOST"),
		inquiryItems: utils.GetEnv("CALLBACK_INQUIRY_ITEMS"),
		inquiryDiscounts: utils.GetEnv("CALLBACK_INQUIRY_DISCOUNTS"),
	}
}