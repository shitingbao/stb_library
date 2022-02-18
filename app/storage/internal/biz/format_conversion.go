package biz

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"stb-library/lib/excel"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type FormatConversionUseCase struct {
	log            *log.Helper
	defaultFileDir DefaultFileDir
}

func NewExportCase(d DefaultFileDir, logger log.Logger) *FormatConversionUseCase {
	return &FormatConversionUseCase{defaultFileDir: d, log: log.NewHelper(logger)}
}

// FileChange file to csv or excel
func (e *FormatConversionUseCase) FileChange(ctx *gin.Context) (string, error) {
	multipartForm, err := ctx.MultipartForm()
	if multipartForm == nil || err != nil {
		return "", errors.New("file form have errr :" + err.Error())
	}
	fHeader, err := ctx.FormFile("file")
	if err != nil {
		return "", err
	}
	fileAdree, err := e.getUpdateFile(fHeader)
	if err != nil {
		return "", err
	}
	partsForm, err := ctx.MultipartForm()
	if err != nil {
		return "", err
	}
	sep, createSep, createFileType, isGBK, isCreateGBK := getFormValues(partsForm.Value)
	resURL := ""
	switch createFileType {
	case "csv":
		fileAdree, err := e.fileToCsv(fileAdree, sep, createSep, isGBK, isCreateGBK)
		if err != nil {
			return "", err
		}
		resURL = fileAdree
	case "excel":
		fileAdree, err := e.fileToExcel(fileAdree, sep, isGBK)
		if err != nil {
			return "", err
		}
		resURL = fileAdree
	}
	return resURL, nil
}

//sep文件分割符，传入文件gbk是否gbk格式，createSep生成文件的分割符， isCreateGBK生成文件是否gbk格式，createFileType 内容为csv或者excel，代表生成哪种文件
func getFormValues(multipartForm map[string][]string) (sep, createSep, createFileType string, isGBK, isCreateGBK bool) {
	for k, v := range multipartForm { //获取表单字段
		switch k {
		case "sep":
			sep = v[0]
		case "gbk":
			if v[0] == "true" {
				isGBK = true
			}
		case "createSep":
			createSep = v[0]
		case "isCreateGBK":
			if v[0] == "true" {
				isCreateGBK = true
			}
		case "createFileType":
			createFileType = v[0]
		}
	}
	return
}

//isGBK true标识使用gbk解析,isCreateGBK标识生成的csv是否用gbk，true代表使用,createSep标识生成文件的间隔符
//只能解析xlsx , csv , txt三种文件，都生成csv
func (e *FormatConversionUseCase) fileToCsv(fileURL, sep, createSep string, isGBK, isCreateGBK bool) (string, error) {
	fileData := [][]string{}
	switch path.Ext(fileURL) {
	case ".xlsx":
		fd, err := excel.ExportParse(fileURL)
		if err != nil {
			return "", err
		}
		fileData = fd
	case ".csv", ".txt":
		fileData = excel.PaseCscOrTxt(fileURL, sep, isGBK)
	default:
		return "", errors.New("file type error")
	}
	fileName := strings.TrimSuffix(path.Base(fileURL), path.Ext(fileURL))
	fileName = strings.Replace(fileName, " ", "-", -1)
	fileName = strings.Replace(fileName, "_", "-", -1)
	fileAdree := path.Join(e.defaultFileDir.DefaultAssetsPath, fileName+".csv")
	switch {
	case isCreateGBK:
		if err := excel.CreateGBKCsvFile(fileAdree, createSep, fileData); err != nil {
			return "", err
		}
	default:
		if err := excel.CreateCsvFile(fileAdree, createSep, fileData); err != nil {
			return "", err
		}
	}
	return fileAdree, nil
}

func (e *FormatConversionUseCase) fileToExcel(fileURL, sep string, isGBK bool) (string, error) {
	ftype := path.Ext(fileURL)
	if ftype != ".csv" && ftype != ".txt" {
		return "", errors.New("file type error")
	}
	fileData := excel.PaseCscOrTxt(fileURL, sep, isGBK)
	fileName := strings.TrimSuffix(path.Base(fileURL), path.Ext(fileURL))
	fileName = strings.Replace(fileName, " ", "-", -1)
	fileName = strings.Replace(fileName, "_", "-", -1)
	fileAdree := path.Join(e.defaultFileDir.DefaultAssetsPath, fileName+".xlsx")
	if err := excel.CreateExcel(fileAdree, fileData); err != nil {
		return "", err
	}
	return fileAdree, nil
}

//获取表单中的文件，保存至默认路径并反馈保存的文件名
func (e *FormatConversionUseCase) getUpdateFile(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	fileAdree := path.Join(e.defaultFileDir.DefaultAssetsPath, file.Filename)
	fl, err := os.Create(fileAdree)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(fl, f); err != nil {
		return "", err
	}
	return fileAdree, nil
}
