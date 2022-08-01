package socketio

import (
	"log"

	socket "github.com/googollee/go-socket.io"
)

var (
	remoteFlag = "remote"

	remoteSuccess = "200"
	remoteFailed  = "500"
)

// 远程模块
func (s *SocketEvent) remoteControl() {
	s.Server.OnConnect("/remote", s.remoteConnectHandle)
	s.Server.OnEvent("/remote", remoteFlag, s.eventRemoteHandle)
	s.Server.OnDisconnect("/remote", s.disconnectRemoteHandle)
	s.Server.OnError("/remote", func(s socket.Conn, e error) {
		log.Println("OnError:", e)
	})

}

func (s *SocketEvent) remoteConnectHandle(con socket.Conn) error {
	log.Println("uid is connect start ======", con.Namespace(), con.Rooms())
	uid, err := s.getUser(con)
	if err != nil {
		return err
	}
	s.loadCon(con, uid)
	return nil
}

// 需要竞争 remote 锁状态，成功反馈 200 代表可连接，否则反馈 500，代表已经有连接
func (s *SocketEvent) eventRemoteHandle(con socket.Conn, msg string) {
	log.Println("con:", con.Namespace(), "--msg:", msg)
	if s.remoteLock.Lock() {
		uid, _ := s.getUser(con)
		s.remoteOwner = uid
		con.Emit(remoteFlag, remoteSuccess)
		return
	}
	uid, _ := s.getUser(con)
	if uid == s.remoteOwner {
		con.Emit(remoteFlag, uid+" :you are remoting")
		return
	}
	con.Emit(remoteFlag, remoteFailed)
}

// 断开连接触发,并清除本地连接,远程锁解锁
func (s *SocketEvent) disconnectRemoteHandle(con socket.Conn, reason string) {
	log.Println("disconnectHandle closed", reason, "=--:", con.Namespace())
	uid, err := s.getUser(con)
	if err != nil {
		return
	}
	s.delCon(uid)
	s.remoteLock.Unlock()
}
