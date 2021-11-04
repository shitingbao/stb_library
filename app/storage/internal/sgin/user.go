package sgin

import (
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) login(ctx *gin.Context) {
	user := &biz.ArgUser{}
	if err := ctx.Bind(user); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	token, err := s.user.Login(ctx, user)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, token)
}

func (s *Sgin) register(ctx *gin.Context) {
	user := &biz.ArgUser{}
	if err := ctx.Bind(user); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	if err := s.user.UserRegister(ctx, user); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, nil)
}
