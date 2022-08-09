package socketio

import (
	"log"

	socket "github.com/googollee/go-socket.io"
)

var (
	chatflag = "chat"
)

func (s *SocketEvent) chatControl() {
	s.Server.OnConnect("/chat", s.chatConnectHandle)
	s.Server.OnEvent("/chat", chatflag, s.eventChatHandle)
	s.Server.OnEvent("/chat", "msg", func(s socket.Conn, msg string) string {
		return "recv " + msg
	})

	s.Server.OnDisconnect("/chat", s.disconnectChatHandle)
	s.Server.OnError("/chat", func(s socket.Conn, e error) {
		log.Println("OnError:", e)
	})
}

// 这里的 err 会直接断开连接,统一管理连接
func (s *SocketEvent) chatConnectHandle(con socket.Conn) error {
	log.Println("uid is connect start ======", con.Namespace(), con.Rooms())
	return nil
}

// 向 room 的 uid 发送一个 200 消息
func (s *SocketEvent) eventChatHandle(con socket.Conn, uid, room string) {
	log.Println("eventChatHandle")
}

func (s *SocketEvent) disconnectChatHandle(con socket.Conn, reason string) {
	log.Println("disconnectChatHandle closed", reason, "=--:", con.Namespace())
}
