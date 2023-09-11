package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/pkg/upload"
)

type FileInfo struct {
	ImageTitle string `json:"imageTitle"`
	ImageUrl   string `json:"imageUrl"`
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	preDirPath := upload.GetSavePreDirPath() + fileName
	if err := upload.SaveFile(fileHeader, uploadSavePath+"/"+preDirPath); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + preDirPath

	return &FileInfo{ImageTitle: "", ImageUrl: accessUrl}, nil
}
