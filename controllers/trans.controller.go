package controller

import (
	"cc-auth/constants"
	"cc-auth/controllers/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller)AddCC(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:http.StatusOK,
	}
	req:=models.ReqCC{}
	if err:=ctx.BindJSON(&req);err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}

	err:=c.usecase.AddCC(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	ctx.JSON(http.StatusOK,res)
}

func (c *controller)TopUpCC(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:http.StatusOK,
	}

	req:=models.ReqCC{}
	if err:=ctx.BindJSON(&req);err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}

	err:=c.usecase.TopUpCC(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	ctx.JSON(http.StatusOK,res)
}