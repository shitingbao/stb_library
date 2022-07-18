package main

import (
	"log"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	router := gin.New()

	server := socketio.NewServer(nil)

	// 连接触发
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("connetc 1234")
		log.Println("connected:", s.ID())
		return nil
	})

	// 事件函数，第一个参数为命名空间，第二个为 room ，用于区分连接和收发消息的标志
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("message", "service notice message")
	})

	// 这里 命名空间不同，需要在路由后的 chat 命名空间地址才能访问
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) {
		// s.SetContext(msg)
		log.Println("chat:", msg)
		s.Emit("message", "service chat message")
		// return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) {
		last := s.Context().(string)
		log.Println("last:", last)
		s.Emit("bye", "this is byte")
		// s.Close()
		// return last
	})

	// 出现 err
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	// 断开连接触发
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	// 连接 url，兼容 get 和 post 两种方法
	router.GET("/socket.io/*any", middle, gin.WrapH(server))
	router.POST("/socket.io/*any", middle, gin.WrapH(server))
	// router.StaticFS("/public", http.Dir("../asset"))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

func middle(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	ctx.Header("Access-Control-Allow-Credentials", "true")
}
