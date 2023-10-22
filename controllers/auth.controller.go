package controller

import (
	"cc-auth/constants"
	"cc-auth/controllers/models"
	auth "cc-auth/controllers/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller) Register(ctx *gin.Context) {
	res := models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:    http.StatusOK,
	}
	req := auth.Credentials{}

	if err := ctx.BindJSON(&req); err != nil {
		res.Message = constants.INVALID_INPUT
		res.Code = http.StatusBadRequest
		ctx.JSON(http.StatusOK, res)
		return
	}

	if err:=c.usecase.Register(req);err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusBadGateway
		ctx.JSON(res.Code,res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *controller)Login(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:http.StatusOK,
	}
	req:=auth.Credentials{}

	if err:=ctx.BindJSON(&req);err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		ctx.JSON(res.Code,res)
		return
	}

	token,err:=c.usecase.Login(req)

	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(res.Code,res)
		return
	}

	res.Data=token
	ctx.JSON(res.Code,res)
}

func (c *controller)DelCC(ctx *gin.Context){
	id:=ctx.Param("id")
	res:=models.GeneralResponse{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	err:=c.usecase.DelCC(id)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusBadRequest
		ctx.JSON(res.Code,res)
		return
	}
	ctx.JSON(res.Code,res)
}