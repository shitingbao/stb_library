package biz

import (
	"errors"
	"io"
	"os"
	"path"
	"stb-library/lib/comparison"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pborman/uuid"
)

type ComparisonUseCase struct {
	log            *log.Helper
	defaultFileDir DefaultFileDir
}

type comparisonParam struct {
	CompareFile comparison.ParisonFileObject
	AimFile     comparison.ParisonFileObject
}

func NewFileComparisonCase(d DefaultFileDir, logger log.Logger) *ComparisonUseCase {
	return &ComparisonUseCase{defaultFileDir: d, log: log.NewHelper(logger)}
}

//post中分api请求比对和表单比对
// api中直接输入文件路径
// 表单中获取文件以及相关文件标识
// 标识参数left，lft，lsep，listitle / right，rft，rsep，ristitle
//分别是两个文件的相关标识（左右）：文件，文件类型，文件分隔标识符，是否是标题
func (f *ComparisonUseCase) FileComparsion(ctx *gin.Context) (comparison.ParisonResult, error) {
	rPath, lPath, err := f.getFormFile(ctx)
	if err != nil {
		return comparison.ParisonResult{}, err
	}
	res, err := comparison.FileComparison(rPath, lPath)
	if err != nil {
		return comparison.ParisonResult{}, err
	}
	return res, nil
}

//根据传入文件名称标识，文件类型标识，从formdata中获取文件
//首字母‘l’代表left，左侧文件，参数说明：lsep 分隔符标识，默认为‘,’ listitle 是否首行是标题，默认为true,首行是标题 lisgbk 当是文本格式文件，标识是否是gbk，默认为utf8
//右侧文件同上
func (f *ComparisonUseCase) getFormFile(ctx *gin.Context) (comparison.ParisonFileObject, comparison.ParisonFileObject, error) {
	leftObject, rightObject := comparison.ParisonFileObject{}, comparison.ParisonFileObject{}
	if ctx.Request.MultipartForm == nil {
		return leftObject, rightObject, errors.New("form is nil")
	}
	for k, v := range ctx.Request.MultipartForm.Value { //获取表单字段
		switch k {
		case "lsep":
			leftObject.Sep = v[0]
		case "listitle":
			if v[0] == "true" {
				leftObject.IsTitle = true
			}
		case "lisgbk":
			if v[0] == "true" {
				leftObject.IsGBK = true
			}
		case "rsep":
			rightObject.Sep = v[0]
		case "ristitle":
			if v[0] == "true" {
				rightObject.IsTitle = true
			}
		case "risgbk":
			if v[0] == "true" {
				rightObject.IsGBK = true
			}
		}
	}
	ladree, err := f.getSaveFilePath(ctx, "left")
	if err != nil {
		return leftObject, rightObject, err
	}
	leftObject.FileName = ladree
	radree, err := f.getSaveFilePath(ctx, "right")
	if err != nil {
		return leftObject, rightObject, err
	}
	rightObject.FileName = radree
	return leftObject, rightObject, nil
}

//获取表单中的文件，保存至默认路径并反馈保存的文件路径
func (f *ComparisonUseCase) getSaveFilePath(ctx *gin.Context, fileName string) (string, error) {
	_, file, err := ctx.Request.FormFile(fileName)
	if err != nil {
		return "", err
	}
	fileHead, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileHead.Close()
	ft := path.Ext(file.Filename)
	if err := os.MkdirAll(f.defaultFileDir.DefaultAssetsPath, os.ModePerm); err != nil {
		return "", err
	}
	fileAdree := path.Join(f.defaultFileDir.DefaultAssetsPath, uuid.NewUUID().String()+ft)
	fl, err := os.Create(fileAdree)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(fl, fileHead); err != nil {
		return "", err
	}
	return fileAdree, nil
}
