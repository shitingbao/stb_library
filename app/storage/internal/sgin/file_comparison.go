package sgin

import (
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) fileComparsion(ctx *gin.Context) {

	res, err := s.comparison.FileComparsion(ctx)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	response.JsonOK(ctx, res)
}
