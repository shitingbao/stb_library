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

	res, err := s.xiaoji.GetMenuList(menu.Id)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, res)
}
