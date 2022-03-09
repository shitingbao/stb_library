package server

import (
	"context"
	"net/http"
	"stb-library/app/storage/internal/conf"
	"stb-library/lib/ws"

	"github.com/gin-gonic/gin"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, g *gin.Engine, h *ws.Hub) *khttp.Server {
	httpSrv := khttp.NewServer(khttp.Address(c.Http.Addr))

	httpSrv.HandleFunc("/sockets/chat", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(context.TODO(), h, w, r)
	})

	httpSrv.HandlePrefix("/", g)

	// v1.RegisterGreeterHTTPServer(httpSrv)
	return httpSrv
}
