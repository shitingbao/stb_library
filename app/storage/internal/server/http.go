package server

import (
	"context"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"stb-library/app/storage/internal/conf"
	"stb-library/lib/ws"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// 自己解析静态资源
func init() {
	mime.AddExtensionType(".js", "text/javascript")
	mime.AddExtensionType(".css", "text/css; charset=utf-8")
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, g *gin.Engine, h *ws.Hub) *khttp.Server {
	var opts = []khttp.ServerOption{
		khttp.Middleware(
			ratelimit.Server(),
		), khttp.Address(c.Http.Addr),
	}
	httpSrv := khttp.NewServer(opts...)
	// 预先处理静态资源
	httpSrv.HandleFunc("/", assetsIndex)
	httpSrv.HandleFunc("_app.config.js", assetsRoute)
	httpSrv.HandleFunc("favicon.ico", assetsRoute)

	httpSrv.HandlePrefix("/assets", http.HandlerFunc(assetsRoute))
	httpSrv.HandlePrefix("/resource", http.HandlerFunc(assetsRoute))

	httpSrv.HandleFunc("/sockets/chat", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(context.TODO(), h, w, r)
	})

	httpSrv.HandlePrefix("/", g)

	// v1.RegisterGreeterHTTPServer(httpSrv)
	return httpSrv
}

// assets 静态资源反馈
func assetsIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("/opt/nginx/dist", r.URL.String(), "index.html"))
	return
}

// assets 静态资源反馈
func assetsRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("/opt/nginx/dist", r.URL.String()))
	return
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
