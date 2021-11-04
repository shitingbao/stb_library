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

func JsonOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  "",
		"data": data,
	})
}

func JsonErr(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		err = errors.New("")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": faile,
		"msg":  err.Error(),
		"data": data,
	})
}
