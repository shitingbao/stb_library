package sgin

import (
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) imageZoom(ctx *gin.Context) {
	list, err := s.image.ImageToWords(ctx)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, list)
}
