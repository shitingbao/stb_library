package sgin

import (
	"net/http"
	"os"
	"path"
	"stb-library/lib/context"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

// 上传视频
func (s *Sgin) fileTransform(ctx *gin.Context) {
	err := s.transform.Transform(context.New(ctx, s.GetUser().UserName))
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx)
}

type fileModel struct {
	FileName string `form:"file_name" json:"file_name"`
}

//TODO Test资源文件下载
func (s *Sgin) downloadFileService(ctx *gin.Context) {

	arg := fileModel{}
	ctx.Bind(&arg)

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+arg.FileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	b, err := os.ReadFile(path.Join(s.defaultFileDir.DefaultAssetsPath, arg.FileName))
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	e := path.Ext(arg.FileName)
	ctx.Data(http.StatusOK, e[1:], b)
	response.JsonOK(ctx)
}
