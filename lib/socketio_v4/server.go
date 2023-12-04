package socketiov4

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/zishang520/socket.io/socket" // 支持 4 以上的其他版本
)

// 跨域中间件
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许跨域访问的域名和端口
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// 继续处理下一个处理函数
		next.ServeHTTP(w, r)
	})
}

func serverLoad() {
	router := http.NewServeMux()
	handler := corsMiddleware(router)
	io := socket.NewServer(nil, nil)
	router.Handle("/socket.io/", io.ServeHandler(nil))
	go http.ListenAndServe(":5005", handler)

	io.On("connection", func(clients ...any) {
		log.Println("connect")
		client := clients[0].(*socket.Socket)
		client.On("heartbeat", func(datas ...any) {
			log.Println("heart")
		})
		client.On("disconnect", func(...any) {
			log.Println("disconnect")
		})
	})

	exit := make(chan struct{})
	SignalC := make(chan os.Signal)

	signal.Notify(SignalC, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range SignalC {
			switch s {
			case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				close(exit)
				return
			}
		}
	}()

	<-exit
	io.Close(nil)
	os.Exit(0)
}

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
func socketioWithGin() {
	g := gin.Default()
	io := socket.NewServer(nil, nil)
	io.Of("/user", nil).On("connection", func(clients ...any) {
		log.Println("connect")
		client := clients[0].(*socket.Socket)
		client.On("ping", func(datas ...any) {
			log.Println("heart")
			client.Emit("pong", "pong")
		})
		client.On("disconnect", func(...any) {
			log.Println("disconnect")
		})
	})
	sock := io.ServeHandler(nil)
	g.Use(cross)
	g.GET("/socket.io/", gin.WrapH(sock))
	g.POST("/socket.io/", gin.WrapH(sock))
	g.Run(":5005")
}
