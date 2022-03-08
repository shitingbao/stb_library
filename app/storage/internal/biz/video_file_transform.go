package biz

import (
	"errors"
	"stb-library/lib/ffmpeg"

	"github.com/gin-gonic/gin"
)

type TransformUseCase struct {
	slog *SlogUseCase

	defaultFileDir DefaultFileDir
}

func NewTransformCase(defaultDir DefaultFileDir, s *SlogUseCase) *TransformUseCase {
	return &TransformUseCase{defaultFileDir: defaultDir, slog: s}
}

// Transform 视频文件类型转换 ftype 参数为完整的文件后缀 .mp4
func (t *TransformUseCase) Transform(ctx *gin.Context) ([]string, error) {
	fileType := ctx.Request.FormValue("ftype")
	if fileType == "" {
		return []string{}, errors.New("file type have nil")
	}
	filePaths, err := getAllFormFile(ctx, t.defaultFileDir.DefaultAssetsPath)
	if err != nil {
		return nil, err
	}
	outFileList := []string{}
	for _, filePath := range filePaths {
		outFilePath, err := t.createTransformFiles(fileType, filePath)
		if err != nil {
			return outFileList, err
		}
		outFileList = append(outFileList, outFilePath)
	}
	return outFileList, nil
}

func (t *TransformUseCase) createTransformFiles(fileType, filePath string) (string, error) {
	fmg := ffmpeg.NewFfmpeg(
		ffmpeg.WithFfmpegOrder(t.defaultFileDir.DefaultDirPath),
		ffmpeg.WithFfmpegTargetType(fileType),
	)
	return fmg.DefaultTransform(filePath)
}
