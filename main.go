package main

import (
	controller "cc-auth/controllers"
	postgre "cc-auth/databases/postgresql"
	redis_db "cc-auth/databases/redis"
	host "cc-auth/hosts"
	callbackHost "cc-auth/hosts/callback"
	router "cc-auth/routers"
	usecase "cc-auth/usecases"
)

func main() {
	merchants:=callbackHost.InitCallback()
	host:=host.InitHost(merchants)
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis,host)
	con := controller.InitController(uc)

	router.MainRouter(con)

}
