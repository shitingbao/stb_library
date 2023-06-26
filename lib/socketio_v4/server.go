package socketiov4

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
