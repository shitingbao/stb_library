package biz

import (
	"errors"
	"stb-library/lib/context"
	"stb-library/lib/ffmpeg"
	"stb-library/lib/ws"
	"time"

	"github.com/sirupsen/logrus"
)

type TransformUseCase struct {
	slog           *SlogUseCase
	hub            *ws.Hub
	defaultFileDir DefaultFileDir
}

func NewTransformCase(defaultDir DefaultFileDir, s *SlogUseCase, h *ws.Hub) *TransformUseCase {
	return &TransformUseCase{
		defaultFileDir: defaultDir,
		slog:           s,
		hub:            h,
	}
}

// Transform 视频文件类型转换 ftype 参数为完整的文件后缀 .mp4
func (t *TransformUseCase) Transform(ctx *context.GContext) error {
	fileType := ctx.Request.FormValue("ftype")
	if fileType == "" {
		return errors.New("file type have nil")
	}
	filePaths, err := getAllFormFile(ctx, t.defaultFileDir.DefaultAssetsPath)
	if err != nil {
		return err
	}
	go func() {

		mes := ws.Message{
			User:     ctx.Username,
			DataType: ws.MessageVideo,
			DateTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		defer func() {
			t.hub.BroadcastUser <- mes
		}()
		outFileList := []string{}
		for _, filePath := range filePaths {
			outFilePath, err := t.createTransformFiles(fileType, filePath)
			if err != nil {
				logrus.Info(err)
				mes.DataType = ws.MessageErr
				mes.Data = err.Error()
				return
			}
			outFileList = append(outFileList, outFilePath)
		}
		mes.Data = outFileList
	}()

	return nil
}

func (t *TransformUseCase) createTransformFiles(fileType, filePath string) (string, error) {
	fmg := ffmpeg.NewFfmpeg(
		ffmpeg.WithFfmpegOrder(t.defaultFileDir.DefaultDirPath),
		ffmpeg.WithFfmpegTargetType(fileType),
	)
	return fmg.DefaultTransform(filePath)
}
