package biz

import (
	"context"
	"net/http"
	"os"
	"path"
	base "stb-library/lib/file_base"
	"stb-library/lib/formopera"

	"github.com/go-kratos/kratos/v2/log"
)

// 默认资源存储路径，当前路径的 assets 下
var DefaultFilePath = ""

func init() {
	dir, err := os.Getwd()
	if err != nil {
	}
	DefaultFilePath = path.Join(dir, "assets")
}

type TransformRepo interface{}

type TransformUseCase struct {
	baseFilePath string
	repo         TransformRepo
	log          *log.Helper
}

func NewTransformUseCase(repo TransformRepo, logger log.Logger) *TransformUseCase {
	return &TransformUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (t *TransformUseCase) Transform(ctx context.Context, r *http.Request) ([]string, error) {
	fileHands := formopera.GetAllFormFiles(r)
	outFileList := []string{}
	for _, fileHand := range fileHands {
		filePath, err := base.SaveFile(DefaultFilePath, fileHand)
		if err != nil {
			return outFileList, nil
		}
		outFileList = append(outFileList, filePath)
	}
	return outFileList, nil
}
