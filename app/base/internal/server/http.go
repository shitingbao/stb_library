package server

import (
	"context"
	"net/http"
	"stb-library/app/base/internal/conf"
	"stb-library/lib/ws"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, g *gin.Engine, h *ws.Hub) *khttp.Server {
	var opts = []khttp.ServerOption{
		khttp.Middleware(
			ratelimit.Server(),
		), khttp.Address(c.Http.Addr),
		khttp.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"http://socket1.cn"}),
			handlers.AllowedHeaders([]string{"Content-Type,AccessToken,X-CSRF-Token, Authorization"}),
			handlers.AllowedMethods([]string{"POST, GET, OPTIONS"}),
			handlers.ExposedHeaders([]string{"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type"}),
			handlers.AllowCredentials(),
		)),
	}
	httpSrv := khttp.NewServer(opts...)

	httpSrv.HandleFunc("/sockets/chat", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(context.TODO(), h, w, r)
	})

	httpSrv.HandlePrefix("/", g)

	// v1.RegisterGreeterHTTPServer(httpSrv)
	return httpSrv
}
