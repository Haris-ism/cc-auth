package grpc_transaction

import (
	"cc-auth/protogen/merchant"
	"cc-auth/utils"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type(
	transactionGrpc struct{
		transcationConn merchant.TransServicesClient
	}
	TransactionInterface interface{
		TransItems(req *merchant.ReqTransItems) (*merchant.ResMerchantTransModel, error)
		// InquiryItems()(*merchant.InquiryMerchantItemsModel,error)
		// InquiryDiscounts()(*merchant.InquiryMerchantDiscountsModel,error)
	}
)

func (g *transactionGrpc)TransItems(req *merchant.ReqTransItems)(*merchant.ResMerchantTransModel,error){
	res,err:=g.transcationConn.TransItems(context.Background(),req)
	if err != nil {
		log.Println("Error on grpc trans :", err)
	}
	return res,nil
}

func InitGrpcTransaction()TransactionInterface{
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(utils.GetEnv("TRANSACTION_HOST_GRPC"),opts...)
	if err!=nil{
		log.Println("failed to dial grpc trans:",err)
	}
	
	transactionConn:=merchant.NewTransServicesClient(conn)
	log.Println("grpc trans connected")
	return &transactionGrpc{
		transcationConn:transactionConn,
	}
}