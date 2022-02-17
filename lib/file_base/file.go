package base

import (
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/pborman/uuid"
)

// SaveFile 传入文件句柄,文件名加上随机字符串片段
func SaveFile(fileDir string, fileHand *multipart.FileHeader) (string, error) {
	baseFile, err := fileHand.Open()
	if err != nil {
		return "", err
	}
	defer baseFile.Close()

	filename := fileHand.Filename
	if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
		return "", err
	}

	fileAdree := path.Join(fileDir, uuid.NewUUID().String()+filename)

	f, err := os.OpenFile(fileAdree, os.O_WRONLY|os.O_CREATE, 0777) //等待拆分
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, baseFile)
	if err != nil {
		return "", err
	}
	return fileAdree, nil
}
