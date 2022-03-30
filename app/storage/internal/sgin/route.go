package sgin

import (
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

var (
	tokenKey = "Authorization"
)

func cross(ctx *gin.Context) {
	// ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Origin", "localhost,http://127.0.0.1,http://124.70.156.31,http://socket1.cn")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Next()
}

func (s *Sgin) verification(ctx *gin.Context) {
	info, err := s.user.GetUserInfo(ctx.GetHeader(tokenKey))
	if err != nil || info.UserName == "" {
		ctx.Abort()
		response.JsonErr(ctx, err, nil)
	}
	s.userInfo = info
	ctx.Next()
	return
}

func (s *Sgin) setRoute() {
	s.g.Use(cross)
	rg := s.g.Group("/api")
	{
		rg.POST("/login", s.login)
		rg.GET("/logout", s.logout)
		rg.POST("/register", s.register)

		rg.GET("/downfile", s.downloadFileService)
		dataRout := rg.Group("/app").Use(s.verification)
		{
			dataRout.GET("/userinfo", s.getUserInfo)
			dataRout.POST("/transform", s.fileTransform)
			dataRout.POST("/qrcode", s.qrcodeDecoder)

			dataRout.POST("/comparison", s.fileComparsion)
			dataRout.POST("/imagezoom", s.imageZoom)

		}
	}

	// s.g.StaticFS("assets", http.Dir(s.defaultFileDir.DefaultAssetsPath))// 直接播放视频
	// s.g.StaticFile("assets", s.defaultFileDir.DefaultAssetsPath)

	s.g.GET("/health", s.health)
	s.g.Static("assets", s.defaultFileDir.DefaultAssetsPath)
}

func (s *Sgin) health(ctx *gin.Context) {
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
