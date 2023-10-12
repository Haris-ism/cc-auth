package controller

import (
	usecase "cc-auth/usecases"

	"github.com/gin-gonic/gin"
)

type (
	controller struct {
		usecase usecase.UsecaseInterface
	}
	ControllerInterface interface {
		Ping(ctx *gin.Context)
		WriteRedis(ctx *gin.Context)
		ReadRedis(ctx *gin.Context)
		InsertPostgre(ctx *gin.Context)
		QueryPostgre(ctx *gin.Context)
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		AddCC(ctx *gin.Context)
		TopUpCC(ctx *gin.Context)
		GetCC(ctx *gin.Context)
		DelCC(ctx *gin.Context)
		TransItem(ctx *gin.Context)
		InquiryItems(ctx *gin.Context)
		InquiryDiscounts(ctx *gin.Context)
	}
)

func InitController(uc usecase.UsecaseInterface) ControllerInterface {
	return &controller{
		usecase: uc,
	}
}
