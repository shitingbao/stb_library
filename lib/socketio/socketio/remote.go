package socketio

import (
	// "IotSocket/rtc"
	"log"
	"stb-library/lib/socketio/rtc"

	socket "github.com/googollee/go-socket.io"
)

var (
	remoteFlag = "remote"
	tokenFlag  = "token"

	remoteSuccess = "200"
	remoteFailed  = "500"
)

// 远程模块
// 开启时，所有待连接的机器连接，并生成对应 room
func (s *SocketEvent) remoteControl() {
	s.Server.OnConnect("/remote", s.remoteConnectHandle)
	s.Server.OnEvent("/remote", "remote", s.eventRemoteHandle)
	s.Server.OnEvent("/remote", "token", s.tokenHandle)
	s.Server.OnDisconnect("/remote", s.disconnectRemoteHandle)
	s.Server.OnError("/remote", func(s socket.Conn, e error) {
		log.Println("OnError:", e)
	})

}

// 	机器连接生成 room，用户连接加入 room
func (s *SocketEvent) remoteConnectHandle(con socket.Conn) error {

	uid := s.getUser(con)
	sn := s.getSn(con)
	log.Println("uid is connect start ======", con.Namespace(), con.Rooms(), "--:uid:", uid, "--sn:", sn)
	s.loadCon(con, uid, sn)
	log.Println("machines:", s.machines, "--roomUsers:", s.roomUsers)
	return nil
}

// 传入想要连接的 sn
// 需要竞争 remote 锁状态，成功反馈 200 代表可连接，否则反馈 500，代表已经有连接
func (s *SocketEvent) eventRemoteHandle(con socket.Conn, sn string) {
	log.Println("con:", con.Namespace(), "--sn:", sn)
	uid := s.getUser(con)
	owner, ok := s.remoteMachines(uid, sn)
	if ok {
		con.Emit(remoteFlag, remoteSuccess)
		return
	}

	if uid == owner { // 判断是否是自己
		con.Emit(remoteFlag, uid+" :you are remoting")
		return
	}
	con.Emit(remoteFlag, remoteFailed)
}

func (s *SocketEvent) tokenHandle(con socket.Conn, channelID string, user string) {
	log.Println("token:", channelID, user)
	token, err := rtc.CreateRtcToken(channelID, user)
	if err != nil {
		con.Emit(tokenFlag, err.Error())
		return
	}
	log.Println("send token:", token)
	con.Emit(tokenFlag, token)
}

// 断开连接触发,并清除本地连接,远程锁解锁
func (s *SocketEvent) disconnectRemoteHandle(con socket.Conn, reason string) {
	log.Println("disconnectHandle closed", reason, "=--:", con.Namespace())
	uid := s.getUser(con)
	sn := s.getSn(con)

	s.delCon(con, uid, sn)
	log.Println("disconnect:", s.machines, "--roomuser:", s.roomUsers)
}
