package socketio

import (
	"log"
	"net/url"
	"stb-library/lib/socketio/rtc"
	"sync"

	socket "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

var (
	snFlag      = "sn" // sn 代表加入拿个房间
	machineFlag = "machine"
	noOwner     = "noOwer"
)

type userCon struct {
	con  socket.Conn
	sn   string
	room string
}
type SocketEvent struct {
	Server        *socket.Server
	roomUsers     map[string][]userCon   // 持久化用户，sn 为 key 的房间分类
	lock          *sync.Mutex            // 控制连接锁，持久化用户连接时可能有 map 异步操作
	machinesOwner map[string]string      // sn 为 key 的所有机器,val string 写入控制者
	machinesCons  map[string]socket.Conn // 保持所有 sn 为 key 的机器连接
}

func New(opts *engineio.Options) *SocketEvent {
	return &SocketEvent{
		Server:        socket.NewServer(opts),
		roomUsers:     make(map[string][]userCon),
		lock:          &sync.Mutex{},
		machinesOwner: make(map[string]string),
		machinesCons:  make(map[string]socket.Conn),
	}
}

func (s *SocketEvent) Control() {
	s.remoteControl()
	s.chatControl()

	s.Server.OnDisconnect("/", func(con socket.Conn, reason string) {
		log.Println("disconnectChatHandle closed", reason, "=--:", con.Namespace())
	})
	s.Server.OnError("/", func(s socket.Conn, e error) {
		log.Println("OnError:", e)
	})

	go func() {
		if err := s.Server.Serve(); err != nil {
			log.Fatalf("socket listen error: %s\n", err)
		}
	}()
}

func (s *SocketEvent) Serve() error {
	return s.Server.Serve()
}

func (s *SocketEvent) Close() {
	s.Server.Close()
}

func (s *SocketEvent) getSn(con socket.Conn) string {
	return getQuery(snFlag, con)
}

func (s *SocketEvent) getMachine(con socket.Conn) string {
	return getQuery(machineFlag, con)
}

func (s *SocketEvent) getWidth(con socket.Conn) string {
	return getQuery("width", con)
}

func (s *SocketEvent) getHeight(con socket.Conn) string {
	return getQuery("height", con)
}

// getQuery 获取 query 参数中的 key 对应的值，默认只取第一位
func getQuery(key string, con socket.Conn) string {
	q, err := url.ParseQuery(con.URL().RawQuery) // 所有socket对象后续获取到的query都可以
	if err != nil {
		return ""
	}
	m, ok := q[key]
	if !ok || len(m) == 0 {
		return ""
	}
	return m[0]
}

// uid 为空的代表机器，否则就是用户连接
func (s *SocketEvent) loadCon(con socket.Conn, sn, msn string) {
	if msn != "" {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.machinesOwner[sn] = noOwner
		s.machinesCons[sn] = con
	}
}

// 将连接加入房间，并获取控制房间的锁所属
// uid 代表 连接者的 sn，第二个 sn代表被连接的机器
// 被连接的 sn 机器就是 room key,加入房间,注意判断已经在房间的状态
// 获取 sn 对应机器控制，反馈拥有者 和 标识 成功 true 失败 false
func (s *SocketEvent) remoteMachines(con socket.Conn, uid, sn string) (string, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.joinRoom(con, uid, sn)
	owner, ok := s.machinesOwner[sn]
	if !ok {
		return "", false
	}
	log.Println(uid, ":--:", sn, "--owner:", owner)
	if owner != noOwner { // 说明已经被获取控制权
		return owner, false
	}
	s.machinesOwner[sn] = uid
	return uid, true
}

func (s *SocketEvent) joinRoom(con socket.Conn, uid, room string) {
	users := s.roomUsers[room]
	for _, user := range users {
		if user.sn == uid {
			return
		}
	}
	s.roomUsers[room] = append(s.roomUsers[room], userCon{
		con:  con,
		sn:   uid,
		room: room,
	})
}

// 一个连接向另一个连接发送消息
func (s *SocketEvent) sendFlagMessage(event, sn string, val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	c := s.machinesCons[sn]
	c.Emit(event, val)

	// cons, ok := s.roomUsers[sn]
	// if !ok {
	// 	return
	// }
	// log.Println("sendFlagMessage===:", event, sn, uid, val)
	// for _, c := range cons {
	// 	if c.sn == uid {
	// 		c.con.Emit(event, val)
	// 	}
	// }
}

// 删除自己拥有的控制，去除房间内的连接保持
// 房间内无连接，关闭rtc 连接
// 注意 con close 是阻塞的
func (s *SocketEvent) delCon(con socket.Conn, sn string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for ksn, owner := range s.machinesOwner {
		if owner == sn {
			s.machinesOwner[ksn] = noOwner
			log.Println("锁定释放：", ksn)
			return
		}
	}
	roomKey := ""
	roomIdx := -1
	flag := false
	for _, users := range s.roomUsers {
		if flag {
			break
		}
		for i, c := range users {
			if c.sn == sn {
				roomKey = c.room
				roomIdx = i
				flag = true
				break
			}
		}
	}
	if roomIdx >= 0 {
		s.roomUsers[roomKey] = append(s.roomUsers[roomKey][:roomIdx], s.roomUsers[roomKey][roomIdx+1:]...)
		log.Println("delete connect==!!!", roomKey, roomIdx)
	}

	if len(s.roomUsers[roomKey]) == 1 {
		rtc.DeleteChannel(roomKey)
		// log.Println("rtc.DeleteChannel(roomKey)==!!!")
	}
}
