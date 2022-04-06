package sgin

import (
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

// MenuList 目录列表
func (s *Sgin) MenuList(ctx *gin.Context) {
	menu := &biz.Menu{}
	if err := ctx.Bind(menu); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	res, err := s.xiaoji.GetMenuList(menu.UserId, menu.ParentId)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, res)
}

func (s *Sgin) CreateMenu(ctx *gin.Context) {
	menu := &biz.Menu{}
	if err := ctx.Bind(menu); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	if err := s.xiaoji.CreateMenu(menu.UserId, menu.ParentId, menu.Name); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, nil)
}

func (s *Sgin) DeleteMenu(ctx *gin.Context) {
	menu := &biz.ArgMenu{}
	if err := ctx.Bind(menu); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	if err := s.xiaoji.DeleteMenu(menu.UserId, menu.Id, menu.ParentId, menu.NewSort); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, nil)
}
