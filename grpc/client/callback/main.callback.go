package grpc_callback

import (
	"cc-auth/protogen/merchant"
	"cc-auth/utils"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type(
	callbackGrpc struct{
		callbackConn merchant.MerchantServicesClient
	}
	CallbackInterface interface{
		InquiryItems()(*merchant.InquiryMerchantItemsModel,error)
		InquiryDiscounts()(*merchant.InquiryMerchantDiscountsModel,error)
	}
)

func (g *callbackGrpc)InquiryItems()(*merchant.InquiryMerchantItemsModel,error){
	res,err:=g.callbackConn.InquiryItems(context.Background(),&emptypb.Empty{})
	if err != nil {
		log.Println("Error on grpc callback :", err)
	}
	return res,nil
}
func (g *callbackGrpc)InquiryDiscounts()(*merchant.InquiryMerchantDiscountsModel,error){
	res,err:=g.callbackConn.InquiryDiscounts(context.Background(),&emptypb.Empty{})
	if err != nil {
		log.Println("Error on grpc callback :", err)
	}
	return res,nil
}

func InitGrpcCallback()CallbackInterface{
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(utils.GetEnv("CALLBACK_HOST_GRPC"),opts...)
	if err!=nil{
		log.Println("failed to dial grpc callback:",err)
	}
	
	callbackConn:=merchant.NewMerchantServicesClient(conn)
	log.Println("grpc callback connected")
	return &callbackGrpc{
		callbackConn:callbackConn,
	}
}