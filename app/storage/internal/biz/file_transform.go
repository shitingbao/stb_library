package biz

import (
	"errors"
	"mime/multipart"
	"stb-library/lib/ffmpeg"
	base "stb-library/lib/file_base"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type TransformUseCase struct {
	defaultFileDir DefaultFileDir
	log            *log.Helper
}

func NewTransformCase(defaultDir DefaultFileDir, logger log.Logger) *TransformUseCase {
	return &TransformUseCase{defaultFileDir: defaultDir, log: log.NewHelper(logger)}
}

// ftype 参数为完整的文件后缀 .txt
func (t *TransformUseCase) Transform(ctx *gin.Context) ([]string, error) {
	fileType := ctx.Request.FormValue("ftype")
	if fileType == "" {
		return []string{}, errors.New("file type have nil")
	}
	formFiles, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	outFileList := []string{}
	fileHandList := []*multipart.FileHeader{}
	for _, fileHands := range formFiles.File {
		fileHandList = append(fileHandList, fileHands...)
	}
	for _, fileHand := range fileHandList {
		filePath, err := base.SaveFile(t.defaultFileDir.DefaultAssetsPath, fileHand)
		if err != nil {
			return outFileList, err
		}
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
