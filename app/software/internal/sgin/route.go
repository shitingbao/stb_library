package sgin

import (
	"mime"
	"net/http"
	"path"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

var (
	tokenKey = "Authorization"
)

var (
	m = minify.New() //资源缩小
)

func init() {
	mime.AddExtensionType(".js", "text/javascript")
	mime.AddExtensionType(".css", "text/css; charset=utf-8")
	m.AddFunc(".js", js.Minify)
	m.AddFunc(".css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
}

func (s *Sgin) setRoute() {
	s.g.Use(cross)

	s.g.StaticFile("/", path.Join(s.defaultFileDir.DefaultDirPath, "dist", "index.html"))             // 指定资源文件 127.0.0.1/ 这种
	s.g.StaticFile("/favicon.ico", path.Join(s.defaultFileDir.DefaultDirPath, "dist", "favicon.ico")) // 127.0.0.1/favicon.ico
	s.g.StaticFile("/_app.config.js", path.Join(s.defaultFileDir.DefaultDirPath, "dist", "_app.config.js"))

	s.g.StaticFS("/assets", http.Dir(path.Join(s.defaultFileDir.DefaultDirPath, "dist", "assets")))     // 以 assets 为前缀的 url
	s.g.StaticFS("/resource", http.Dir(path.Join(s.defaultFileDir.DefaultDirPath, "dist", "resource"))) // 比如 127.0.0.1/resource

	// s.g.StaticFS("assets", http.Dir(s.defaultFileDir.DefaultDirPath))// 直接播放视频
	// s.g.StaticFile("assets", s.defaultFileDir.DefaultDirPath)

	codeRoute := s.g.Group("/code")
	{
		codeRoute.POST("/create", s.codeCreate)
		codeRoute.POST("/list", s.codeList)
		codeRoute.POST("/docx", s.getDocx)

	}

	codeHeaderRoute := s.g.Group("/code_header")
	{
		codeHeaderRoute.POST("/create", s.createHeaders)
		codeHeaderRoute.POST("/list", s.codeHeaderList)

	}

	// s.g.Static("assets", s.defaultFileDir.DefaultAssetsPath)
}

// option 过滤
func cross(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	// ctx.Header("Access-Control-Allow-Origin", "http://socket1.cn")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	//允许类型校验
	if ctx.Request.Method == "OPTIONS" {
		ctx.JSON(http.StatusOK, "ok")
		return
	}
	ctx.Next()
}
