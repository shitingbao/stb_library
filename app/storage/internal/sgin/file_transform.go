package sgin

import (
	"net/http"
	"os"
	"path"
	"stb-library/lib/response"

	"github.com/gin-gonic/gin"
)

func (s *Sgin) upload(ctx *gin.Context) {
	list, err := s.transform.Transform(ctx)
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}
	response.JsonOK(ctx, list)
}

type fileModel struct {
	FileName string `form:"file_name" json:"file_name"`
}

//TODO Test资源文件下载
func (s *Sgin) downloadFileService(ctx *gin.Context) {

	arg := fileModel{}
	ctx.Bind(&arg)

	ctx.Header("Content-Type", "video/mpeg4")
	ctx.Header("Content-Disposition", "attachment; filename="+path.Join(s.defaultFileDir.DefaultAssetsPath, arg.FileName))
	ctx.Header("Content-Transfer-Encoding", "binary")
	b, err := os.ReadFile(path.Join(s.defaultFileDir.DefaultAssetsPath, arg.FileName))
	if err != nil {
		response.JsonErr(ctx, err, nil)
		return
	}

	ctx.Data(http.StatusOK, "", b)
	response.JsonOK(ctx)
}
