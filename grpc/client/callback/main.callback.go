package grpc_callback

import (
	"cc-auth/protogen/merchant"
	"cc-auth/utils"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type(
	callbackGrpc struct{
		callbackConn merchant.InquiryServicesClient
		// callbackHost		string
		// inquiryItems		string
		// inquiryDiscounts	string
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
	fmt.Println("inquiry")
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
		fmt.Println("failed to dial grpc callback:",err)
	}
	
	callbackConn:=merchant.NewInquiryServicesClient(conn)
	fmt.Println("grpc callback connected")
	return &callbackGrpc{
		callbackConn:callbackConn,
		// callbackHost:utils.GetEnv("CALLBACK_HOST"),
		// inquiryItems: utils.GetEnv("CALLBACK_INQUIRY_ITEMS"),
		// inquiryDiscounts: utils.GetEnv("CALLBACK_INQUIRY_DISCOUNTS"),
	}
}