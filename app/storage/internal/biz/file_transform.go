package biz

import (
	"context"
	"net/http"
	base "stb-library/lib/file_base"
	"stb-library/lib/formopera"

	"github.com/go-kratos/kratos/v2/log"
)

type TransformUseCase struct {
	defaultFileDir DefaultFileDir
	log            *log.Helper
}

func NewTransformCase(defaultDir DefaultFileDir, logger log.Logger) *TransformUseCase {
	return &TransformUseCase{defaultFileDir: defaultDir, log: log.NewHelper(logger)}
}

func (t *TransformUseCase) Transform(ctx context.Context, r *http.Request) ([]string, error) {
	fileHands := formopera.GetAllFormFiles(r)
	outFileList := []string{}
	for _, fileHand := range fileHands {
		filePath, err := base.SaveFile(t.defaultFileDir.DefaultFilePath, fileHand)
		if err != nil {
			return outFileList, err
		}
		outFileList = append(outFileList, filePath)
	}
	return outFileList, nil
}
