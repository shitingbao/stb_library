package sgin

import (
	"stb-library/app/storage/internal/biz"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type Sgin struct {
	bg   *biz.GreeterUsecase
	user *biz.UserUseCase
	log  *log.Helper
	g    *gin.Engine
}

// sgin 只作路由对应
func NewSgin(b *biz.GreeterUsecase, u *biz.UserUseCase, logger log.Logger) *gin.Engine {
	ginModel := gin.Default()
	s := &Sgin{
		bg:   b,
		user: u,
		log:  log.NewHelper(logger),
		g:    ginModel,
	}
	s.setRoute()
	return s.g
}

func (s *Sgin) setRoute() {
	rg := s.g.Group("/basic-api")
	{
		rg.POST("/login", s.login)
		rg.GET("/logout", s.logout)
		rg.POST("/register", s.register)

		rg.GET("/userinfo", s.getUserInfo)

	}
}
