package sgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) helloworld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "ok",
		"data": "hello!",
	})
}
