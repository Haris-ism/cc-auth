package transaction

import (
	"cc-auth/constants"
	"cc-auth/utils"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type(
	transaction struct{
		transactionHost		string
		transactionItems	string
	}
	TransactionInterface interface{
		Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error)
	}
)

func (h *transaction)Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error){
	var url string
	var method string
	switch types{
		case constants.TRANSACTION_ITEMS:
			url=h.transactionHost+h.transactionItems
			method=constants.HTTP_POST
	}
	res,data,err:=utils.HTTPRequest(url,method,body,header)
	
	if err!=nil{
		return res,data,err
	}

	return res,data,nil
}

func InitTransaction()TransactionInterface{
	return &transaction{
		transactionHost:utils.GetEnv("TRANSACTION_HOST"),
		transactionItems: utils.GetEnv("TRANSACTION_ITEMS"),
	}
}