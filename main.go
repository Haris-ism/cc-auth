package main

import (
	controller "cc-auth/controllers"
	postgre "cc-auth/databases/postgresql"
	redis_db "cc-auth/databases/redis"
	host "cc-auth/hosts"
	callbackHost "cc-auth/hosts/callback"
	"cc-auth/hosts/transaction"
	router "cc-auth/routers"
	usecase "cc-auth/usecases"
)

func main() {
	callback:=callbackHost.InitCallback()
	transaction:=transaction.InitTransaction()
	host:=host.InitHost(callback,transaction)
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis,host)
	con := controller.InitController(uc)

	router.MainRouter(con)

}
