package sgin

import (
	"net/http"
	"stb-library/app/storage/internal/biz"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) helloworld(ctx *gin.Context) {
	m := &biz.Greeter{}
	if err := ctx.Bind(m); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  err.Error(),
		})
		return
	}
	da, err := s.bg.SayHello(ctx, m)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "ok",
		"data": da,
	})
}
