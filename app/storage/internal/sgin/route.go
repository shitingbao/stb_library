package sgin

import (
	"context"
	"net/http"
	"os"
	v1 "stb-library/api/storage/v1"
	"stb-library/app/storage/internal/biz"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type Sgin struct {
	v1.UnimplementedStorageServer
	user      *biz.UserUseCase
	transform *biz.TransformUseCase
	center    *biz.CentralUseCase
	log       *log.Helper
	g         *gin.Engine
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

// sgin 只作路由对应
func NewSgin(ginModel *gin.Engine, u *biz.UserUseCase, c *biz.CentralUseCase, logger log.Logger) *Sgin {
	s := &Sgin{
		user:   u,
		center: c,
		log:    log.NewHelper(logger),
		g:      ginModel,
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	s.setRoute(dir)
	return s
}

func (s *Sgin) setRoute(dir string) {
	rg := s.g.Group("/basic-api")
	{
		rg.POST("/login", s.login)
		rg.GET("/logout", s.logout)
		rg.POST("/register", s.register)

		rg.GET("/userinfo", s.getUserInfo)
		rg.POST("/upload", s.upload)

	}

	s.g.StaticFS("assets", http.Dir(dir))
}

func (s *Sgin) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{
		Message: req.GetName(),
	}, nil
}
