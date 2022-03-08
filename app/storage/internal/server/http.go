package server

import (
	"stb-library/app/storage/internal/conf"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, g *gin.Engine) *http.Server {
	httpSrv := http.NewServer(http.Address(c.Http.Addr))
	httpSrv.HandlePrefix("/", g)

	// v1.RegisterGreeterHTTPServer(httpSrv)
	return httpSrv
}
