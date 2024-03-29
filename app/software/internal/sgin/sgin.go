package sgin

import (
	"context"
	"encoding/json"
	"os"
	"path"
	v1 "stb-library/api/software/v1"
	"stb-library/app/software/internal/biz"
	"stb-library/lib/ws"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"net/http"
	_ "net/http/pprof"
)

const (
	formFileSize = 65
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewGinEngine,
	NewSgin,
	ConstructorDefaultDir,
	NewChatSocketfunc,
)

type Sgin struct {
	v1.UnimplementedSoftwareServer

	g              *gin.Engine
	defaultFileDir biz.DefaultFileDir
	code           *biz.CodeUseCase
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

// ConstructorDefaultDir 默认当前路径下放资源目录
func ConstructorDefaultDir() (biz.DefaultFileDir, error) {
	dir, err := os.Getwd()
	if err != nil {
		return biz.DefaultFileDir{}, err
	}
	defaultDir := biz.DefaultFileDir{
		DefaultAssetsPath: path.Join(dir, "assets"),
		DefaultDirPath:    dir,
	}

	if err := os.MkdirAll(defaultDir.DefaultAssetsPath, os.ModePerm); err != nil {
		return defaultDir, err
	}
	return defaultDir, nil
}

// sgin 只作路由对应
func NewSgin(dir biz.DefaultFileDir, ginModel *gin.Engine, codeCase *biz.CodeUseCase) *Sgin {
	ginModel.MaxMultipartMemory = formFileSize << 20 // 为了 form 提交文件做前提
	s := &Sgin{
		g:              ginModel,
		defaultFileDir: dir,
		code:           codeCase,
	}
	s.setRoute()

	go http.ListenAndServe("127.0.0.1:6060", nil) // 增加 pprof 检查入口

	return s
}

func NewChatSocketfunc() *ws.Hub {
	h := ws.NewHub(func(data []byte, hub *ws.Hub) error {
		msg := ws.Message{}
		if err := json.Unmarshal(data, &msg); err != nil {
			return err
		}
		//原样消息发公告
		hub.Broadcast <- msg
		return nil
	})
	go h.Run(context.TODO())
	return h
}
