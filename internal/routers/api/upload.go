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

	params := &service.ClientUploadParams{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, params)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToResponse(code.ErrorInvalidParams.WithDetails(errs.Errors()...))
		return
	}

	var svcUploadFileData *service.FileInfo
	var svc = service.New(c.Request.Context())
	var err error

	file, fileHeader, errf := c.Request.FormFile("imagefile")

	if errf == nil {
		defer file.Close()
		svcUploadFileData, err = svc.UploadFile(upload.TypeImage, file, fileHeader, params)
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
