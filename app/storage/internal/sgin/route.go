package sgin

import (
	"log"
	"mime"
	"net/url"
	"path/filepath"
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/response"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	tokenKey = "Authorization"
)

// 自己解析静态资源
func init() {
	mime.AddExtensionType(".js", "text/javascript")
	mime.AddExtensionType(".css", "text/css; charset=utf-8")
}

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
	s.g.Group("/").Use(s.assets)
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

			dataRout.GET("/ghealth", s.health)
			dataRout.POST("/phealth", s.health)
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

// assets 静态资源反馈
func (s *Sgin) assets(ctx *gin.Context) {
	log.Println("URL=====:", ctx.Request.URL.String())
	//反馈静态主页，需要下面css，和js以及fonts的资源路径配合
	//这里的第二个参数就是你要反馈的资源地址，也就是类似index.html的完整路径
	switch ctx.Request.URL.String() {
	case "/":
		s.g.StaticFile("/opt/nginx", filepath.Join("dist", ctx.Request.URL.String(), "index.html"))
		return
	case "/_app.config.js", "favicon.ico":
		s.g.StaticFile("/opt/nginx", filepath.Join("dist", ctx.Request.URL.String()))
		return
	}

	paths, err := parsePaths(ctx.Request.URL)
	log.Println("parsePaths======:", paths, err)
	//这里的path反馈工作元素内容待定
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(nil)
		return
	}

	if paths[0] == "assets" || paths[0] == "resource" {
		s.g.StaticFile("/opt/nginx", filepath.Join("dist", ctx.Request.URL.String()))
		// http.ServeFile(w, r, filepath.Join(str, "dist", paths[0], paths[len(paths)-1]))
		return
	}

}

//解析url
func parsePaths(u *url.URL) ([]string, error) {
	paths := []string{}
	pstr := u.EscapedPath()
	for _, str := range strings.Split(pstr, "/")[1:] {
		s, err := url.PathUnescape(str)
		if err != nil {
			return nil, err
		}
		paths = append(paths, s)
	}
	return paths, nil
}
