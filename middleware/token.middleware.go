package middleware

import (
	"cc-auth/constants"
	"cc-auth/controllers/models"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func TokenAuth(ctx *gin.Context){
	res:=models.GeneralResponse{}
	ReqHeader:=models.ReqHeader{}
	err:=ctx.BindHeader(&ReqHeader)
	if err!=nil{
		res.Message=constants.ERROR_TOKEN
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		ctx.Abort()
		return
	}
	if strings.HasPrefix(ReqHeader.Authorization, "Bearer ") == false {
		res.Message="Invalid Token"
		res.Code=http.StatusForbidden
		ctx.JSON(http.StatusForbidden, res)
		ctx.Abort()
		return
	}
	tokenString:=strings.Replace(ReqHeader.Authorization, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			res.Message="Token Expired"
			res.Code=http.StatusForbidden
			ctx.JSON(http.StatusForbidden, res)
			ctx.Abort()
			return
		}
	} else {
		fmt.Println(err, reflect.TypeOf(err), reflect.ValueOf(err).Kind())
		res.Message="Token Error"
		res.Code=http.StatusForbidden
		ctx.JSON(http.StatusForbidden, res)
		ctx.Abort()
		return
	}
	ctx.Next()
}