package sgin

import (
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) qrcodeDecoder(ctx *gin.Context) {
	list, err := s.qrcode.QrcodeDecoder(ctx)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, list)
}
