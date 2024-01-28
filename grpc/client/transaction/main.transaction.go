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
		transcationConn merchant.MerchantServicesClient
	}
	TransactionInterface interface{
		TransItems(ctx context.Context, req *merchant.ReqTransItemsModel)(*merchant.ResMerchantTransModel,error)
	}
)

func (g *transactionGrpc)TransItems(ctx context.Context, req *merchant.ReqTransItemsModel)(*merchant.ResMerchantTransModel,error){
	res:=&merchant.ResMerchantTransModel{}
	res,err:=g.transcationConn.TransItems(ctx,req)
	if err != nil {
		log.Println("Error on grpc trans :", err)
		return res,err
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
	
	transactionConn:=merchant.NewMerchantServicesClient(conn)
	log.Println("grpc trans connected")

	return &transactionGrpc{
		transcationConn:transactionConn,
	}
}