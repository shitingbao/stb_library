package biz

import (
	"stb-library/lib/formopera"
	"stb-library/lib/images"
	imagetowordapi "stb-library/lib/imagetowordAPI"

	"time"

	"github.com/gin-gonic/gin"
)

//ImageWord 业务类
type ImageWordUseCase struct {
	slog            *SlogUseCase
	accessToken     string
	accessTokenDate string
	baidubceAddress string

	defaultFileDir DefaultFileDir
}

func NewImageToWordCase(defaultDir DefaultFileDir, s *SlogUseCase) *ImageWordUseCase {
	return &ImageWordUseCase{defaultFileDir: defaultDir, slog: s}
}

//Post 图片转文字
func (im *ImageWordUseCase) ImageToWords(ctx *gin.Context) ([]imagetowordapi.AcceptResultWord, error) {
	imageURLs, err := im.getFileHands(ctx)
	if err != nil {

		return nil, err
	}

	result, err := im.imagesOpera(imageURLs)
	if err != nil {

		return nil, err
	}

	return result, nil
}

//获取表单内图片保存，并反馈对应所有图片路径
func (im *ImageWordUseCase) getFileHands(ctx *gin.Context) ([]string, error) {
	imageURLs := []string{}
	fileHands := formopera.GetAllFormFiles(ctx.Request)
	for _, v := range fileHands {
		file, err := v.Open()
		if err != nil {
			return imageURLs, err
		}
		imageURL, err := images.ByteToImage(im.defaultFileDir.DefaultAssetsPath, file)
		if err != nil {
			return imageURLs, err
		}
		imageURLs = append(imageURLs, imageURL)
		file.Close()
	}
	return imageURLs, nil
}

//imagesOpera 传入图片路径，亲求三方接口反馈文字对象,需要先检查token可用性
func (im *ImageWordUseCase) imagesOpera(imageURLs []string) ([]imagetowordapi.AcceptResultWord, error) {
	result := []imagetowordapi.AcceptResultWord{}
	token, err := imagetowordapi.CheckTokenEffect(im.accessTokenDate)
	if err != nil {
		return result, err
	}
	if token != "" {
		im.accessToken = token
		im.accessTokenDate = time.Now().Format("2006-01-02 15:04:05")
		// core.WebConfig.SaveConfig()
	}

	for _, v := range imageURLs {
		imagesBase64 := []string{}
		base64, err := images.ImageToBase64(v)
		if err != nil {
			return result, err
		}
		imagesBase64 = append(imagesBase64, base64)
		res, err := imagetowordapi.GetImageWord(im.baidubceAddress, im.accessToken, im.accessTokenDate, imagesBase64)
		if err != nil {
			return result, err
		}
		result = append(result, res)
	}
	return result, nil
}
