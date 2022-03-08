package sgin

import (
	"os"
	"path"
	v1 "stb-library/api/storage/v1"
	"stb-library/app/storage/internal/biz"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewGinEngine,
	NewSgin,
	ConstructorDefaultDir,
)

type Sgin struct {
	v1.UnimplementedStorageServer
	center           *biz.CentralUseCase
	formatConversion *biz.FormatConversionUseCase
	comparison       *biz.ComparisonUseCase
	transform        *biz.TransformUseCase
	image            *biz.ImageWordUseCase
	qrcode           *biz.QrcodeUseCase
	user             *biz.UserUseCase

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
) *Sgin {
	ginModel.MaxMultipartMemory = 20 << 20 // 为了 form 提交文件做前提

	s := &Sgin{
		center:           c,
		comparison:       cmp,
		transform:        trans,
		formatConversion: ex,
		image:            img,
		qrcode:           q,
		user:             u,
		g:                ginModel,
		defaultFileDir:   dir,
	}
	s.setRoute()
	return s
}
