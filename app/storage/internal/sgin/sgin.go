package sgin

import (
	"context"
	"encoding/json"
	"os"
	"path"
	v1 "stb-library/api/storage/v1"
	"stb-library/app/storage/internal/biz"
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
	v1.UnimplementedStorageServer
	userInfo         *biz.UserResult
	center           *biz.CentralUseCase
	formatConversion *biz.FormatConversionUseCase
	comparison       *biz.ComparisonUseCase
	transform        *biz.TransformUseCase
	image            *biz.ImageWordUseCase
	qrcode           *biz.QrcodeUseCase
	user             *biz.UserUseCase
	xiaoji           *biz.XiaojiUseCase
	imgZoom          *biz.ImageZoomUseCase

	g              *gin.Engine
	defaultFileDir biz.DefaultFileDir
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
func NewSgin(dir biz.DefaultFileDir, ginModel *gin.Engine,
	ex *biz.FormatConversionUseCase, cmp *biz.ComparisonUseCase, trans *biz.TransformUseCase,
	img *biz.ImageWordUseCase, q *biz.QrcodeUseCase, u *biz.UserUseCase, c *biz.CentralUseCase,
	imgzoom *biz.ImageZoomUseCase, xiaoji *biz.XiaojiUseCase,
) *Sgin {
	ginModel.MaxMultipartMemory = formFileSize << 20 // 为了 form 提交文件做前提

	s := &Sgin{
		center:           c,
		comparison:       cmp,
		transform:        trans,
		formatConversion: ex,
		image:            img,
		qrcode:           q,
		imgZoom:          imgzoom,
		xiaoji:           xiaoji,
		user:             u,
		g:                ginModel,

		defaultFileDir: dir,
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

func (s *Sgin) GetUser() *biz.UserResult {
	return s.userInfo
}
