package sgin

import (
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func cross(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Origin", "http://127.0.0.1,http://124.70.156.31,http://socket1.cn")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()

}

func (s *Sgin) setRoute() {
	s.g.Use(cross)
	rg := s.g.Group("/api")
	{
		rg.POST("/login", s.login)
		rg.GET("/logout", s.logout)
		rg.POST("/register", s.register)

		rg.GET("/userinfo", s.getUserInfo)
		rg.POST("/transform", s.fileTransform)
		rg.POST("/qrcode", s.qrcodeDecoder)

		rg.POST("/comparison", s.fileComparsion)

		rg.GET("/central", s.sayHello)
		rg.GET("/downfile", s.downloadFileService)
	}

	// s.g.StaticFS("assets", http.Dir(s.defaultFileDir.DefaultAssetsPath))// 直接播放视频
	// s.g.StaticFile("assets", s.defaultFileDir.DefaultAssetsPath)

	s.g.Static("assets", s.defaultFileDir.DefaultAssetsPath)
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
