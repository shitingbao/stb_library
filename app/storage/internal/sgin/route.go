package sgin

import (
	"stb-library/app/storage/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type Sgin struct {
	gServer *service.GreeterService
	log     *log.Helper
	g       *gin.Engine
}

// sgin 只作路由对应
func NewSgin(gs *service.GreeterService, logger *log.Helper) *Sgin {
	ginModel := gin.Default()
	s := &Sgin{
		gServer: gs,
		log:     logger,
		g:       ginModel,
	}
	s.setRoute(ginModel)
	return s
}

func (s *Sgin) setRoute(g *gin.Engine) {
	rg := g.Group("/stb").Use()
	{
		rg.GET("/h", s.helloworld)
	}
}

func (s *Sgin) GetGin() *gin.Engine {
	return s.g
}
