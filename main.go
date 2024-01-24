package main

import (
	controller "cc-auth/controllers"
	postgre "cc-auth/databases/postgresql"
	redis_db "cc-auth/databases/redis"
	grpcClient "cc-auth/grpc/client"
	grpcCallback "cc-auth/grpc/client/callback"
	grpcTransaction "cc-auth/grpc/client/transaction"
	host "cc-auth/hosts"
	"cc-auth/hosts/callback"
	"cc-auth/hosts/transaction"
	router "cc-auth/routers"
	usecase "cc-auth/usecases"
)

func main() {
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()

	callback:=callback.InitCallback()
	transaction:=transaction.InitTransaction()
	callbackGrpc:=grpcCallback.InitGrpcCallback()
	transactionGrpc:=grpcTransaction.InitGrpcTransaction()
	hostGrpc:=grpcClient.InitGrpcClient(callbackGrpc,transactionGrpc)
	host:=host.InitHost(callback,transaction)
	uc := usecase.InitUsecase(postgre, redis,host,hostGrpc)
	con := controller.InitController(uc)

	// go func (){
	// 	res,_:=hostGrpc.Callback().InquiryItems()
	// 	fmt.Println("grpc res:",res)

	// }()
	router.MainRouter(con)

}
