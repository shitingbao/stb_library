package main

import (
	"log"
	"time"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func main() {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = "user"
	opts.Query["pwd"] = "pass"
	uri := "http://127.0.0.1:8000/socket.io/chat"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})
	client.On("bye", func(msg string) {
		log.Printf("bye" + msg)
	})
	for {

		// if err := client.Emit("notice", "stb client notice"); err != nil {
		// 	log.Println("notice err:", err)
		// 	return
		// }

		if err := client.Emit("msg", "stb client msg"); err != nil {
			log.Println("bye err:", err)
			return
		}
		log.Printf("send message")
		time.Sleep(time.Second)
	}
}
