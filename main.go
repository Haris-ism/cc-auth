package main

import (
	controller "cc-auth/controllers"
	postgre "cc-auth/databases/postgresql"
	redis_db "cc-auth/databases/redis"
	router "cc-auth/routers"
	usecase "cc-auth/usecases"
)

func main() {
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis)
	con := controller.InitController(uc)

	router.MainRouter(con)

}
