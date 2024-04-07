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

func (u Upload) Upload(c *gin.Context) {

	var svcUploadFileData *service.FileInfo
	var svc = service.New(c.Request.Context())
	var err error

	response := app.NewResponse(c)

	name := c.Request.FormValue("name")

	if name != "" {
		svcUploadFileData, err = svc.UploadFileByURL(upload.TypeImage, name)
	} else {
		file, fileHeader, errf := c.Request.FormFile("imagefile")
		defer file.Close()

		// dump.P(file)

		if errf != nil {
			response.ToResponse(code.ErrorInvalidParams.WithDetails(errf.Error()))
			return
		}
		svcUploadFileData, err = svc.UploadFile(upload.TypeImage, file, fileHeader)
	}

	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToResponse(code.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(code.Success.WithData(svcUploadFileData))

	return

}
