package middleware

import (
	"cc-auth/constants"
	"cc-auth/controllers/models"
	"cc-auth/utils"
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignatureValidation(ctx *gin.Context){
	res:=models.GeneralResponse{}
	ReqHeader:=models.ReqHeader{}
	err:=ctx.BindHeader(&ReqHeader)
	if err!=nil{
		logrus.Error("err:",err)
		res.Message=constants.ERROR_TOKEN
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		ctx.Abort()
		return
	}
}

func Signature(req string,ts string)string{
	key:=utils.GetEnv("SIG_KEY")
	data:=req+"&"+ts+"&"+key
	// fmt.Println("data:",data)
	res:=HashSha512(key,data)
	// fmt.Println("hash:",res)
	return res
}
func HashSha512(secret, data string) string {
	hash := hmac.New(sha512.New, []byte(secret))
	hash.Write([]byte(data))
	return fmt.Sprintf("%x",hash.Sum(nil))
}