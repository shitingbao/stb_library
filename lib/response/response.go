package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonOK(ctx *gin.Context, err error, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  err.Error(),
		"data": data,
	})
}

func JsonErr(ctx *gin.Context, err error, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10001,
		"msg":  err.Error(),
		"data": data,
	})
}
