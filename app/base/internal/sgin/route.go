package sgin

import (
	"mime"
	"net/http"
	"regexp"
	"stb-library/lib/response"

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

	vueAssetsRoutePath = "/opt/nginx/dist" // dist 所在路径
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
	// s.g.Use(cross)

	// s.g.StaticFile("/", path.Join(vueAssetsRoutePath, "index.html"))             // 指定资源文件 127.0.0.1/ 这种
	// s.g.StaticFile("/favicon.ico", path.Join(vueAssetsRoutePath, "favicon.ico")) // 127.0.0.1/favicon.ico
	// s.g.StaticFile("/_app.config.js", path.Join(vueAssetsRoutePath, "_app.config.js"))

	// s.g.StaticFS("/assets", http.Dir(path.Join(vueAssetsRoutePath, "assets")))     // 以 assets 为前缀的 url
	// s.g.StaticFS("/resource", http.Dir(path.Join(vueAssetsRoutePath, "resource"))) // 比如 127.0.0.1/resource

	// s.g.StaticFS("assets", http.Dir(s.defaultFileDir.DefaultAssetsPath))// 直接播放视频
	// s.g.StaticFile("assets", s.defaultFileDir.DefaultAssetsPath)

	s.g.GET("/health", s.health)
	// s.g.Static("assets", s.defaultFileDir.DefaultAssetsPath)
}

func (s *Sgin) health(ctx *gin.Context) {
	name := ""
	if err := ctx.Bind(name); err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	n, err := s.center.SayHello(name)

	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, n)
}

// option 过滤

func cross(ctx *gin.Context) {
	allowedOrigins := []string{"http://192.168.31.33:3001", "http://192.168.31.33:3000"}
	origin := ctx.Request.Header.Get("Origin")
	// log.Println("origin=:", origin, " Referer:", ctx.Request.Referer()) origin or Referer
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			break
		}
	}
	// ctx.Header("Access-Control-Allow-Origin", "http://localhost:3001,http://localhost:3000")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,x-device-sn,x-device-token")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,x-device-sn")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if ctx.Request.Method == "OPTIONS" {
		ctx.JSON(http.StatusOK, "ok")
		return
	}
	ctx.Next()
}

func (s *Sgin) verification(ctx *gin.Context) {
	// info, err := s.user.GetUserInfo(ctx.GetHeader(tokenKey))
	// if err != nil || info.UserName == "" {
	// 	response.JsonErr(ctx, err, nil)
	// 	ctx.Abort()
	// }
	// s.userInfo = info
	// ctx.Next()
}
