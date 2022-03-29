package biz

import (
	"mime/multipart"
	"stb-library/lib/context"
	base "stb-library/lib/file_base"
)

// 获取所有表单内的文件，并反馈所有文件路径
// 传入基本路径
func getAllFormFile(ctx *context.GContext, basePath string) ([]string, error) {
	formFiles, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	filePaths := []string{}
	fileHandList := []*multipart.FileHeader{}
	for _, fileHands := range formFiles.File {
		fileHandList = append(fileHandList, fileHands...)
	}
	for _, fileHand := range fileHandList {
		filePath, err := base.SaveFile(basePath, fileHand)
		if err != nil {
			return filePaths, err
		}
		filePaths = append(filePaths, filePath)
	}
	return filePaths, nil
}
