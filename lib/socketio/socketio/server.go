package socketio

import (
	// "IotSocket/rtc"
	"log"
	"net/url"
	"stb-library/lib/socketio/rtc"
	"sync"

	socket "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

var (
	usrFlag = "uid" // 需要连接时带上身份标识
	snFlag  = "sn"  // sn 代表加入拿个房间
	noOwner = "noOwer"
)

// type machine struct {
// 	remoteLock  *lock.Locker // remote 已连接锁定，获取该锁代表已经有远程连接，断开后释放
// 	remoteOwner string       // 代表锁的所属用户
// }

type userCon struct {
	con socket.Conn
	uid string
}
type SocketEvent struct {
	Server    *socket.Server
	roomUsers map[string][]userCon // 持久化用户，sn 为 key 的房间分类
	lock      *sync.Mutex          // 控制连接锁，持久化用户连接时可能有 map 异步操作
	machines  map[string]string    //sn 为 key 的所有机器,val string 写入控制者
}

func New(opts *engineio.Options) *SocketEvent {
	return &SocketEvent{
		Server:    socket.NewServer(opts),
		roomUsers: make(map[string][]userCon),
		lock:      &sync.Mutex{},
		machines:  make(map[string]string),
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

func (s *SocketEvent) getUser(con socket.Conn) string {
	return getQuery(usrFlag, con)
}

func (s *SocketEvent) getSn(con socket.Conn) string {
	return getQuery(snFlag, con)
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
func (s *SocketEvent) loadCon(con socket.Conn, uid, sn string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if uid == "" {
		s.machines[sn] = noOwner
	} else {
		s.roomUsers[sn] = append(s.roomUsers[sn], userCon{
			con: con,
			uid: uid,
		})
	}
}

// 获取 sn 对应机器控制，反馈拥有者 和 标识 成功 true 失败 false
func (s *SocketEvent) remoteMachines(uid, sn string) (string, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	owner := s.machines[sn]
	if owner != noOwner {
		return owner, false
	}
	s.machines[sn] = uid
	log.Println("machines:", s.machines)
	return uid, true
}

// 一个连接向另一个连接发送消息
func (s *SocketEvent) sendFlagMessage(con socket.Conn, event, room, uid, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	cons := s.roomUsers[room]
	for _, c := range cons {
		if c.uid == uid {
			c.con.Emit(event, val)
		}
	}
}

// 删除自己拥有的控制，去除房间内的连接保持
// 房间内无连接，关闭rtc 连接
// 注意 con close 是阻塞的
func (s *SocketEvent) delCon(con socket.Conn, uid, sn string) {
	log.Println("start delcon:", uid, sn)
	s.lock.Lock()
	defer s.lock.Unlock()
	if uid == "" {
		return
	}
	owners := s.machines[sn]
	if owners == uid {
		s.machines[sn] = noOwner
		log.Println("锁定释放：", sn)
	}

	for i, c := range s.roomUsers[sn] {
		if c.con == con {
			s.roomUsers[sn] = append(s.roomUsers[sn][:i], s.roomUsers[sn][i+1:]...)
			log.Println("清除连接：", uid, "--sn:", sn)
			continue
		}
	}
	if len(s.roomUsers[sn]) == 1 {
		rtc.DeleteChannel(sn)
	}
}
