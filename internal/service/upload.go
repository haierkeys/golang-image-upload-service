package service

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/pkg/oss"
	"github.com/haierspi/golang-image-upload-service/pkg/upload"
)

type FileInfo struct {
	ImageTitle string `json:"imageTitle"`
	ImageUrl   string `json:"imageUrl"`
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	return svc.fileSyncHandle(fileType, file, fileHeader)
}

func (svc *Service) UploadFileByURL(fileType upload.FileType, url string) (*FileInfo, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	file, err := os.Create(uuid.New().String() + upload.GetFileExt(url))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	muFile, fileHeader, err := upload.FileToMultipart(file)

	if err != nil {
		return nil, err
	}

	return svc.fileSyncHandle(fileType, muFile, fileHeader)

}

func (svc *Service) fileSyncHandle(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {

	var accessUrlPre string
	var fileName string

	// 通过剪切板上传的附件 都是一个默认名字
	if fileHeader.Filename == "image.png" {
		fileName = upload.GetFileName(uuid.New().String() + fileHeader.Filename)
	} else {
		fileName = upload.GetFileName(fileHeader.Filename)
	}

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

	dateDirFileName := upload.GetSavePreDirPath() + fileName
	if err := upload.SaveFile(fileHeader, uploadSavePath+"/"+dateDirFileName); err != nil {
		return nil, err
	}
	accessUrlPre = global.AppSetting.UploadServerUrl

	// 阿里云oss
	if global.OSSSetting.Enable {
		err := oss.UploadByFile(dateDirFileName, file)
		if err != nil {
			return nil, errors.Wrap(err, "oss.UploadByFile err")
		}
	}

	accessUrl := accessUrlPre + "/" + dateDirFileName

	return &FileInfo{ImageTitle: "", ImageUrl: accessUrl}, nil
}
