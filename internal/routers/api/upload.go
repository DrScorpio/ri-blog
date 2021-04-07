package api

import (
	"github.com/DrScorpio/ri-blog/global"
	"github.com/DrScorpio/ri-blog/internal/service"
	"github.com/DrScorpio/ri-blog/pkg/app"
	"github.com/DrScorpio/ri-blog/pkg/convert"
	"github.com/DrScorpio/ri-blog/pkg/errcode"
	"github.com/DrScorpio/ri-blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errResp := errcode.InvaildParams.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvaildParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
