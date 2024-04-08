package api

import (
	"github.com/gin-gonic/gin"

	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/internal/service"
	"github.com/haierspi/golang-image-upload-service/pkg/app"
	"github.com/haierspi/golang-image-upload-service/pkg/code"
	"github.com/haierspi/golang-image-upload-service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// Upload 上传文件
func (u Upload) Upload(c *gin.Context) {

	var svcUploadFileData *service.FileInfo
	var svc = service.New(c.Request.Context())
	var err error
	var name string

	response := app.NewResponse(c)

	file, fileHeader, errf := c.Request.FormFile("imagefile")
	if errf == nil {
		defer file.Close()
		svcUploadFileData, err = svc.UploadFile(upload.TypeImage, file, fileHeader)
	} else if name = c.Request.FormValue("name"); name != "" {
		svcUploadFileData, err = svc.UploadFileByURL(upload.TypeImage, name)
	} else {
		global.Logger.Errorf(c, "app.ErrorInvalidParams len 0:")
		response.ToResponse(code.ErrorInvalidParams)
		return
	}

	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToResponse(code.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(code.Success.WithData(svcUploadFileData))

	return

}
