package sgin

import (
	"stb-library/app/software/internal/biz"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) codeHeaderCreate(ctx *gin.Context) {
	codes := []biz.CodeHeader{}
	if err := ctx.Bind(&codes); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	if err := s.codeHeader.Create(codes); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, nil)
}

func (s *Sgin) codeHeaderGetCodes(ctx *gin.Context) {
	arg := biz.ArgCode{}
	if err := ctx.Bind(&arg); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	res, err := s.codeHeader.GetCodes(arg.Num, arg.Key, arg.Filters)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, res)
}
