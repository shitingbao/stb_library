package main

import (
	// "IotSocket/socketio"
	"log"
	"stb-library/lib/socketio/socketio"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	server := socketio.New(nil)
	server.Control()
	defer server.Close()

	// 连接 url，兼容 get 和 post 两种方法
	router.GET("/socket.io/*any", middle, gin.WrapH(server.Server))
	router.POST("/socket.io/*any", middle, gin.WrapH(server.Server))
	// router.StaticFS("/public", http.Dir("../asset"))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

func middle(ctx *gin.Context) {

	ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("origin"))
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")
}
