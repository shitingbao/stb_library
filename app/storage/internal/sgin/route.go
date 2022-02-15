package sgin

import (
	"net/http"
	"os"
	"path"
	v1 "stb-library/api/storage/v1"
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type Sgin struct {
	v1.UnimplementedStorageServer
	center           *biz.CentralUseCase
	formatConversion *biz.FormatConversionUseCase
	comparison       *biz.ComparisonUseCase
	transform        *biz.TransformUseCase
	image            *biz.ImageWordUseCase
	qrcode           *biz.QrcodeUseCase
	user             *biz.UserUseCase

	log *log.Helper
	g   *gin.Engine
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

// ConstructorDefaultDir 默认当前路径下放资源目录
func ConstructorDefaultDir() (biz.DefaultFileDir, error) {
	defaultDir := biz.DefaultFileDir{
		DefaultFilePath: "./assets",
	}

	if err := os.MkdirAll(defaultDir.DefaultFilePath, os.ModePerm); err != nil {
		return defaultDir, err
	}
	return defaultDir, nil
}

// sgin 只作路由对应
func NewSgin(ginModel *gin.Engine, logger log.Logger,
	ex *biz.FormatConversionUseCase, cmp *biz.ComparisonUseCase, trans *biz.TransformUseCase,
	img *biz.ImageWordUseCase, q *biz.QrcodeUseCase, u *biz.UserUseCase, c *biz.CentralUseCase,
) *Sgin {
	ginModel.MaxMultipartMemory = 20 << 20 // 为了 form 提交文件做前提

	s := &Sgin{
		center:           c,
		comparison:       cmp,
		transform:        trans,
		formatConversion: ex,
		image:            img,
		qrcode:           q,
		user:             u,
		log:              log.NewHelper(logger),
		g:                ginModel,
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	s.setRoute(dir)
	return s
}

func cross(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Origin", "http://127.0.0.1,http://124.70.156.31,http://socket1.cn")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()

}

func (s *Sgin) sayHello(ctx *gin.Context) {
	user := &biz.UserResult{}
	if err := ctx.Bind(user); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	n, err := s.center.SayHello(user.UserName)

	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, n)
}

func (s *Sgin) setRoute(dir string) {
	s.g.Use(cross)
	rg := s.g.Group("/api")
	{
		rg.POST("/login", s.login)
		rg.GET("/logout", s.logout)
		rg.POST("/register", s.register)

		rg.GET("/userinfo", s.getUserInfo)
		rg.POST("/upload", s.upload)

		rg.GET("/central", s.sayHello)
	}

	s.g.StaticFS("assets", http.Dir(path.Join(dir, "assets")))
}
