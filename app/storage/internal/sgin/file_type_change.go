package sgin

import (
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) Formatconversion(ctx *gin.Context) {

	fPath, err := s.formatConversion.FileChange(ctx)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	response.JsonOK(ctx, map[string]string{"fpath": fPath})
}
