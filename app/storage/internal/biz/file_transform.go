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

// DefaultFilePath 默认资源存储路径，当前路径的 assets 下
const DefaultFileDir = "assets"

type TransformRepo interface{}

type TransformUseCase struct {
	baseFilePath string
	repo         TransformRepo
	log          *log.Helper
}

func NewTransformUseCase(repo TransformRepo, logger log.Logger) *TransformUseCase {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	basePath := path.Join(dir, DefaultFileDir)
	return &TransformUseCase{baseFilePath: basePath, repo: repo, log: log.NewHelper(logger)}
}

func (t *TransformUseCase) Transform(ctx context.Context, r *http.Request) ([]string, error) {
	fileHands := formopera.GetAllFormFiles(r)
	outFileList := []string{}
	for _, fileHand := range fileHands {
		filePath, err := base.SaveFile(t.baseFilePath, fileHand)
		if err != nil {
			return outFileList, err
		}
		outFileList = append(outFileList, filePath)
	}
	return outFileList, nil
}
