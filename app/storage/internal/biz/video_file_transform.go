package biz

import (
	"errors"
	"mime/multipart"
	"stb-library/lib/context"
	"stb-library/lib/ffmpeg"
	"stb-library/lib/ws"
	"time"
)

type TransformUseCase struct {
	sLog           *SlogUseCase
	hub            *ws.Hub
	defaultFileDir DefaultFileDir
}

func NewTransformCase(defaultDir DefaultFileDir, s *SlogUseCase, h *ws.Hub) *TransformUseCase {
	return &TransformUseCase{
		defaultFileDir: defaultDir,
		sLog:           s,
		hub:            h,
	}
}

// Transform 视频文件类型转换 ftype 参数为完整的文件后缀 .mp4
func (t *TransformUseCase) Transform(ctx *context.GContext) error {

	fileType := ctx.Request.FormValue("ftype")
	if fileType == "" {
		return errors.New("file type have nil")
	}

	formFiles, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	go t.asncFormOption(ctx.Username, fileType, formFiles)

	return nil
}

// asncFormOption 异步处理文件
func (t *TransformUseCase) asncFormOption(username, fileType string, formFiles *multipart.Form) {
	mes := ws.Message{
		User:     username,
		DataType: ws.MessageVideo,
		DateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	defer func() {
		t.hub.BroadcastUser <- mes
	}()
	filePaths, err := formOption(formFiles, t.defaultFileDir.DefaultAssetsPath)
	if err != nil {
		mes.DataType = ws.MessageErr
		mes.Data = err.Error()
		return
	}
	outFileList := []string{}
	for _, filePath := range filePaths {
		outFilePath, err := t.createTransformFiles(fileType, filePath)
		if err != nil {
			mes.DataType = ws.MessageErr
			mes.Data = err.Error()
			return
		}
		outFileList = append(outFileList, outFilePath)
	}
	mes.Data = outFileList
}

func (t *TransformUseCase) createTransformFiles(fileType, filePath string) (string, error) {
	fmg := ffmpeg.NewFfmpeg(
		ffmpeg.WithFfmpegOrder(t.defaultFileDir.DefaultDirPath),
		ffmpeg.WithFfmpegTargetType(fileType),
	)
	return fmg.DefaultTransform(filePath)
}
