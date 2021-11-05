package sgin

import (
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) logout(ctx *gin.Context) {
	user := &biz.UserResult{}
	if err := ctx.Bind(user); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	s.user.LoginOut(ctx, user.Token)
	response.JsonOK(ctx, nil)
}

func (s *Sgin) login(ctx *gin.Context) {
	user := &biz.ArgUser{}
	if err := ctx.Bind(user); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	userModel, err := s.user.Login(ctx, user)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, userModel)
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
