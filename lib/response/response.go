package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"data": data,
	})
}

func JsonErr(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		err = errors.New("")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10001,
		"msg":  err.Error(),
		"data": data,
	})
}
