package server

import (
	"context"
	"fmt"
	v1 "stb-library/api/storage/v1"
	"stb-library/app/storage/internal/conf"
	"stb-library/app/storage/internal/service"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	// var opts = []http.ServerOption{
	// 	http.Middleware(
	// 		recovery.Recovery(),
	// 	),
	// }
	// if c.Http.Network != "" {
	// 	opts = append(opts, http.Network(c.Http.Network))
	// }
	// if c.Http.Addr != "" {
	// 	opts = append(opts, http.Address(c.Http.Addr))
	// }
	// if c.Http.Timeout != nil {
	// 	opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	// }
	// srv := http.NewServer(opts...)
	srv := ginServer(c)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func ginServer(c *conf.Server) *http.Server {
	router := gin.Default()
	// 使用kratos中间件
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	router.GET("/helloworld/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if name == "error" {
			// 返回kratos error
			kgin.Error(ctx, errors.Unauthorized("auth_error", "no authentication"))
		} else {
			ctx.JSON(200, map[string]string{"welcome": name})
		}
	})

	httpSrv := http.NewServer(http.Address(":8000"))
	httpSrv.HandlePrefix("/", router)
	return httpSrv
}
