package socketio

import (
	"log"
	"stb-library/lib/socketio/rtc"

	socket "github.com/googollee/go-socket.io"
)

// 远程模块
// 开启时，所有待连接的机器连接，并生成对应 room
func (s *SocketEvent) remoteControl() {

	s.Server.OnConnect("/remote", s.remoteConnectHandle)
	s.Server.OnEvent("/remote", "openrtc", s.eventRemoteHandle)

	s.Server.OnEvent("/remote", "keydown", s.eventKeydownHandle)
	s.Server.OnEvent("/remote", "keyup", s.eventKeyupHandle)
	s.Server.OnEvent("/remote", "mousemove", s.eventMousemoveHandle)
	s.Server.OnEvent("/remote", "mousedown", s.eventMousedownHandle)
	s.Server.OnEvent("/remote", "mouseup", s.eventMouseupHandle)

	s.Server.OnDisconnect("/remote", s.disconnectRemoteHandle)
	s.Server.OnError("/remote", func(s socket.Conn, e error) {
		log.Println("OnError:", e)
	})
}

func (s *SocketEvent) eventKeydownHandle(
	con socket.Conn,
	msg struct {
		Code string
		Sn   string
	}) {
	s.sendFlagMessage("keydown", msg.Sn, msg.Code)
}

func (s *SocketEvent) eventKeyupHandle(
	con socket.Conn,
	msg struct {
		Code string
		Sn   string
	}) {
	s.sendFlagMessage("keyup", msg.Sn, msg.Code)
}

func (s *SocketEvent) eventMousemoveHandle(con socket.Conn, msg struct {
	Sn string
	X  string
	Y  string
}) {
	s.sendFlagMessage("mousemove", msg.Sn, struct {
		X string
		Y string
	}{
		X: msg.X,
		Y: msg.Y,
	})
}

func (s *SocketEvent) eventMousedownHandle(con socket.Conn, msg struct {
	Sn  string
	X   string
	Y   string
	Btn string
}) {
	s.sendFlagMessage("mousedown", msg.Sn, struct {
		X   string
		Y   string
		Btn string
	}{
		X:   msg.X,
		Y:   msg.Y,
		Btn: msg.Btn,
	})
}

func (s *SocketEvent) eventMouseupHandle(con socket.Conn, msg struct {
	Sn  string
	X   string
	Y   string
	Btn string
}) {
	s.sendFlagMessage("mouseup", msg.Sn, struct {
		X   string
		Y   string
		Btn string
	}{
		X:   msg.X,
		Y:   msg.Y,
		Btn: msg.Btn,
	})
}

// 	机器连接生成 room，用户连接加入 room
func (s *SocketEvent) remoteConnectHandle(con socket.Conn) error {
	uid := s.getSn(con)
	mSn := s.getMachine(con)
	s.loadCon(con, uid, mSn)
	log.Println(s.machinesCons, s.machinesOwner, s.roomUsers, "--", uid, ":", mSn)
	return nil
}

// 传入想要连接的 sn
// 需要竞争 remote 锁状态，成功反馈 200 代表可连接，否则反馈 500，代表已经有连接
func (s *SocketEvent) eventRemoteHandle(con socket.Conn, msg struct{ Sn string }) {
	log.Println("eventRemoteHandle===:", con.Namespace(), "--sn:", msg.Sn)
	uid := s.getSn(con)
	_, ok := s.remoteMachines(con, uid, msg.Sn)
	if ok {
		//获取rtc token (发送给设备端)
		token, err := rtc.CreateRtcToken(msg.Sn, msg.Sn)
		if err != nil {
			log.Println("CreateRtcToken:", err)
		}
		s.sendFlagMessage("openrtc", msg.Sn, token)
	} else {
		con.Emit("openrtc", "500") // 反馈一个获取控制权失败的信号
	}

	con.Emit("screen", struct {
		Width  string
		Height string
	}{
		Width:  s.getWidth(con),
		Height: s.getHeight(con),
	})
}

// 断开连接触发,并清除本地连接,远程锁解锁
func (s *SocketEvent) disconnectRemoteHandle(con socket.Conn, reason string) {
	sn := s.getSn(con)
	s.delCon(con, sn)
	log.Println("disconnect sn:", s.machinesCons, s.machinesOwner, s.roomUsers)
}
