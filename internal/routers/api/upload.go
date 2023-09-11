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

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("imagefile")

	defer file.Close()

	// dump.P(file)

	if err != nil {
		response.ToResponse(code.ErrorInvalidParams.WithDetails(err.Error()))
		return
	}

	svc := service.New(c.Request.Context())
	svcUploadFileData, err := svc.UploadFile(upload.TypeImage, file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToResponse(code.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(code.Success.WithData(svcUploadFileData))
	return

}
