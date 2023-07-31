package stboutserver

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"stb-library/lib/external_service/core"
	"stb-library/lib/external_service/stbserver"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

var (
	// 设置一个重启的参数，用于区分正常开启还是升级
	argReload      = flag.Bool("reload", false, "listen on fd open 3 (internal use only)")
	defaultAddress = ":4399"
)

type EndlessTcp struct {
	address string
	listen  net.Listener
	wg      *sync.WaitGroup
	*stbserver.UnimplementedStbServerServer
}

func (e *EndlessTcp) HeartBeat(cli stbserver.StbServer_HeartBeatServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	e.wg.Add(1)
	log.Println("add 1")
	sid := ""
	for {
		res, err := cli.Recv()
		if err != nil {
			core.UserHub.DeleteData(sid)
			e.wg.Done()
			log.Println("Done beat")
			return err
		}
		sid = res.Id
		core.UserHub.PutData(sid)
		log.Println("reload2:", res.Id)
	}
}

func NewEndlessTcp() *EndlessTcp {
	return &EndlessTcp{
		address: defaultAddress,
		wg:      &sync.WaitGroup{},
	}
}

// 注意，服务器只能配置一个 UnaryInterceptor和StreamClientInterceptor，
// 否则会报错，客户端也是，虽然不会报错，但是只有最后一个才起作用。
// 如果你想配置多个，可以使用拦截器链，如go-grpc-middleware，或者自己实现。
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("start")                             // 打印开始
	log.Println("info.FullMethod=", info.FullMethod) // 执行了哪个方法，比如执行了 /stbserver.StbServer/GetSummonerInfo
	log.Println("req===:", req)                      // 参数内容，比如 idcard:"qwer"  name:"shitingbao"

	r, ok := req.(*stbserver.Identity) // 进行拦截判断
	if ok && r.Name == "shitingbao" {
		return nil, errors.New("not this name")
	}
	m, err := handler(ctx, req) // 执行实际方法的逻辑

	// 记录请求参数 耗时 错误信息等数据
	// err 获取执行的逻辑错误
	log.Println(err)
	log.Println("end") // 执行完逻辑打印
	return m, err
}

// 流拦截器
// handler 方法执行就是对应 grpc 的方法执行逻辑
// 比如 HeartBeat 方法被调用，先执行 log start 然后执行 HeartBeat ，执行结束后执行 log end ，类似外面包装了一个方法
func streamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("Start of streaming RPC - Method: %s, Timestamp: %s", info.FullMethod, time.Now().Format(time.RFC3339))
	err := handler(srv, ss)
	log.Printf("End of streaming RPC - Method: %s, Timestamp: %s", info.FullMethod, time.Now().Format(time.RFC3339))
	return err
}

func (e *EndlessTcp) EndlessTcpRegisterAndListen() error {
	flag.Parse()
	add, err := net.ResolveTCPAddr("tcp4", e.address)
	if err != nil {
		return err
	}
	if *argReload {
		// 获取到cmd中的ExtraFiles内的文件信息，以它为内容启动监听
		// ExtraFiles的内容在reload方法中放入
		log.Println("EndlessTcpRegisterAndListen reload:", *argReload)
		f := os.NewFile(3, "")
		l, err := net.FileListener(f)
		if err != nil {
			log.Println("FileListener:", err)
			return err
		}
		e.listen = l
	} else {
		l, err := net.ListenTCP("tcp", add)
		if err != nil {
			log.Println("ListenTCP:", err)
			return err
		}
		e.listen = l
	}

	go e.serverLoad()
	e.signalHandler()
	return nil
}

func (e *EndlessTcp) serverLoad() {
	opts := []grpc.ServerOption{}
	// 加上拦截器
	opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor))
	opts = append(opts, grpc.StreamInterceptor(streamServerInterceptor))
	s := grpc.NewServer(opts...)
	stbserver.RegisterStbServerServer(s, e)

	log.Println("start listen:", e.address)
	s.Serve(e.listen)
}

// 信号处理
func (e *EndlessTcp) signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			signal.Stop(ch)
			log.Printf("stop listen")
			return
		case syscall.SIGUSR2:
			if err := e.reload(); err != nil {
				log.Fatalf("restart error: %v", err)
			}
			log.Println("start waiting")
			e.wg.Wait()
			log.Println("stop old listen")
			return
		}
	}
}

func (e *EndlessTcp) reload() error {
	f, err := e.listen.(*net.TCPListener).File()
	if err != nil {
		log.Println("reload", err)
		return err
	}
	cmd := exec.Command(os.Args[0], "-reload")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.ExtraFiles = append(cmd.ExtraFiles, f)
	return cmd.Start()
}
