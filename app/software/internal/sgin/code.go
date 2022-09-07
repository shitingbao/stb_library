package sgin

import (
	"stb-library/app/software/internal/model"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) codeCreate(ctx *gin.Context) {
	arg := model.ArgCodeModel{}
	if err := ctx.Bind(&arg); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	if err := s.code.Create(arg.Codes); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, nil)
}

func (s *Sgin) codeList(ctx *gin.Context) {
	arg := model.ArgDocx{}
	if err := ctx.Bind(&arg); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	res, err := s.code.GetCodes(arg.ContentsNum, arg.Language, arg.ContentsKey, arg.ContentFilters)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, res)
}

func (s *Sgin) createHeaders(ctx *gin.Context) {
	codes := model.ArgCodeModel{}
	if err := ctx.Bind(&codes); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	if err := s.code.CreateHeaders(codes.Codes); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, nil)
}

func (s *Sgin) codeHeaderList(ctx *gin.Context) {
	arg := model.ArgDocx{}
	if err := ctx.Bind(&arg); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	res, err := s.code.GetHeaderCode(arg.ContentsNum, arg.Language, arg.HeaderFilters)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, res)
}

func (s *Sgin) getDocx(ctx *gin.Context) {
	arg := model.ArgDocx{}
	if err := ctx.Bind(&arg); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	res, err := s.code.CreateDocx(arg)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, res)
}
