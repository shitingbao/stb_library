package socketio

import (
	"errors"
	"log"
	"net/url"
	"stb-library/lib/socketio/lock"
	"sync"

	socket "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

var (
	usrFlag = "uid" // 需要连接时带上身份标识
)

type SocketEvent struct {
	Server      *socket.Server
	conMap      map[string]socket.Conn // 持久化用户连接，key 为连接时的 room 标识
	lock        *sync.Mutex            // 控制连接锁
	remoteLock  *lock.Locker           // remote 已连接锁定，锁定代表已经有远程连接
	remoteOwner string
}

func New(opts *engineio.Options) *SocketEvent {
	return &SocketEvent{
		Server:     socket.NewServer(opts),
		conMap:     make(map[string]socket.Conn),
		lock:       &sync.Mutex{},
		remoteLock: lock.NewSLocker(),
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

func (s *SocketEvent) getUser(con socket.Conn) (string, error) {
	return getQuery(usrFlag, con)
}

// getQuery 获取 query 参数中的 key 对应的值，默认只取第一位
func getQuery(key string, con socket.Conn) (string, error) {
	q, err := url.ParseQuery(con.URL().RawQuery) // 所有socket对象后续获取到的query都可以
	if err != nil {
		return "", err
	}
	m, ok := q[key]
	if !ok || len(m) == 0 {
		return "", errors.New(key + " is nil")
	}
	return m[0], nil
}

func (s *SocketEvent) loadCon(con socket.Conn, uid string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.conMap[uid] = con
	// log.Println("uid is connect:", uid)
}

// 一个连接向另一个连接发送消息
func (s *SocketEvent) sendFlagMessage(con socket.Conn, event, uid, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	flagCon, ok := s.conMap[uid]
	if !ok {
		con.Emit(event, val)
		return
	}
	flagCon.Emit(event, val)
	con.Emit(event, val)
}

func (s *SocketEvent) delCon(uid string) {
	s.lock.Lock()
	delete(s.conMap, uid)
	s.lock.Unlock()
}
