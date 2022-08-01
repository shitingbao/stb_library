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
	uid, err := s.getUser(con)
	if err != nil {
		return err
	}
	s.loadCon(con, uid)
	return nil
}

func (s *SocketEvent) eventChatHandle(con socket.Conn, uid string) {
	// log.Println("con:", con.Namespace(), "--uid:", uid)
	s.sendFlagMessage(con, chatflag, uid, "200")
}

func (s *SocketEvent) disconnectChatHandle(con socket.Conn, reason string) {
	log.Println("disconnectChatHandle closed", reason, "=--:", con.Namespace())
	uid, err := s.getUser(con)
	if err != nil {
		return
	}
	s.delCon(uid)
}
