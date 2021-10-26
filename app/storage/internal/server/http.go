package server

import (
	v1 "stb-library/api/storage/v1"
	"stb-library/app/storage/internal/conf"
	"stb-library/app/storage/internal/service"
	"stb-library/app/storage/internal/sgin"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, sg sgin.Sgin, greeter *service.GreeterService, logger log.Logger) *http.Server {
	httpSrv := http.NewServer(http.Address(c.Http.Addr))
	httpSrv.HandlePrefix("/", sg.GetGin())

	v1.RegisterGreeterHTTPServer(httpSrv, greeter)
	return httpSrv
}
