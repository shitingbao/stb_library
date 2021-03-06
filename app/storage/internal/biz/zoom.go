package biz

import (
	"errors"
	"stb-library/lib/context"
	"stb-library/lib/images"
	"strconv"
)

type ImageZoomUseCase struct {
	sLog *SlogUseCase

	defaultFileDir DefaultFileDir
}

func NewImageZoomCase(defaultDir DefaultFileDir, s *SlogUseCase) *ImageZoomUseCase {
	return &ImageZoomUseCase{defaultFileDir: defaultDir, sLog: s}
}

func getFormValue(ctx *context.GContext, content string) (int, error) {
	val := ctx.Request.FormValue(content)
	return strconv.Atoi(val)
}

func getOptionParam(ctx *context.GContext) (int, int, int, error) {

	width, err := getFormValue(ctx, "width")
	if err != nil {
		return 0, 0, 0, errors.New("width have nil")
	}

	height, err := getFormValue(ctx, "height")
	if err != nil {
		return 0, 0, 0, errors.New("height have nil")
	}

	quality, err := getFormValue(ctx, "quality")
	if err != nil {
		return 0, 0, 0, errors.New("quality have nil")
	}
	return width, height, quality, nil
}

// Transform 视频文件类型转换 ftype 参数为完整的文件后缀 .mp4
func (t *ImageZoomUseCase) ImageZoom(ctx *context.GContext) ([]string, error) {
	height, width, quality, err := getOptionParam(ctx)
	if err != nil {
		return nil, err
	}
	img := images.NewImages(
		images.WithHeight(uint(height)),
		images.WithWidth(uint(width)),
		images.WithQuality(quality),
	)

	filePaths, err := getAllFormFile(ctx, t.defaultFileDir.DefaultAssetsPath)
	if err != nil {
		return nil, err
	}

	urlList := []string{}
	for _, ul := range filePaths {
		u, err := img.ImageZoom(ul, t.defaultFileDir.DefaultAssetsPath)
		if err != nil {
			return nil, err
		}
		urlList = append(urlList, u)
	}

	return urlList, nil
}
