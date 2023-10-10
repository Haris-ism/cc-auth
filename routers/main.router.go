package router

import (
	controller "cc-auth/controllers"
	"cc-auth/middleware"
	"cc-auth/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func MainRouter(con controller.ControllerInterface) {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware)
	v1 := r.Group("v1")
	v1.GET("/ping", con.Ping)
	v1.POST("/writeredis", con.WriteRedis)
	v1.POST("/readredis", con.ReadRedis)
	v1.POST("/postgre/insert", con.InsertPostgre)
	v1.GET("/postgre/query", con.QueryPostgre)

	v2 := r.Group("v2")
	v2.POST("/register", con.Register)
	v2.POST("/login", con.Login)
	v2.POST("/add-credit-cards",middleware.TokenAuth,con.AddCC)
	v2.POST("/top-up-credit-cards",middleware.TokenAuth,con.TopUpCC)
	v2.GET("/credit-cards",middleware.TokenAuth,con.GetCC)
	v2.DELETE("/delete-credit-cards/:id",middleware.TokenAuth,con.DelCC)

	logrus.Info("starts")
	r.Run(utils.GetEnv("PORT"))
}
