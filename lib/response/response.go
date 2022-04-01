package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	success = iota
	faile
)

func JsonOK(ctx *gin.Context, data ...interface{}) {

	ctx.Header("Access-Control-Allow-Origin", "http://socket1.cn:8000")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	var res interface{}
	if len(data) > 0 {
		res = data[0]
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  "",
		"data": res,
	})
}

func JsonErr(ctx *gin.Context, err error, data interface{}) {

	ctx.Header("Access-Control-Allow-Origin", "http://socket1.cn:8000")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if err == nil {
		err = errors.New("")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": faile,
		"msg":  err.Error(),
		"data": data,
	})
}
