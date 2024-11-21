package service

import (
    "io"
    "mime/multipart"
    "net/http"
    "os"

    "github.com/google/uuid"
    "github.com/pkg/errors"

    "github.com/haierspi/golang-image-upload-service/global"
    "github.com/haierspi/golang-image-upload-service/pkg/aws_s3"
    "github.com/haierspi/golang-image-upload-service/pkg/cloudflare_r2"
    "github.com/haierspi/golang-image-upload-service/pkg/local_fs"
    "github.com/haierspi/golang-image-upload-service/pkg/oss"
    "github.com/haierspi/golang-image-upload-service/pkg/upload"
)

type FileInfo struct {
    ImageTitle string `json:"imageTitle"`
    ImageUrl   string `json:"imageUrl"`
}

type Uploader interface {
    SendFile(pathKey string, f multipart.File, h *multipart.FileHeader) (string, error)
    SendContent(pathKey string, content []byte) (string, error)
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

    uploadTempPath := upload.GetTempPath() + "/uploads"

    uploadTempFile := uploadTempPath + "/" + uuid.New().String() + upload.GetFileExt(url)

    if upload.CheckPath(uploadTempPath) {
        if err := upload.CreatePath(uploadTempPath, os.ModePerm); err != nil {
            return nil, errors.New("failed to create save directory.")
        }
    }
    if upload.CheckPermission(uploadTempPath) {
        return nil, errors.New("insufficient file permissions.")
    }

    file, err := os.Create(uploadTempFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    defer os.Remove(uploadTempFile)

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        return nil, err
    }

    muFile, fileHeader, err := upload.FileToMultipart(file)

    return svc.fileSyncHandle(fileType, muFile, fileHeader)

}

func (svc *Service) fileSyncHandle(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {

    var fileName string

    // dump.P(fileHeader)

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

    fileKey := upload.GetSavePreDirPath() + fileName

    var up = make(map[string]Uploader)

    var dstFileKey string

    for _, v := range []string{"local_fs", "oss", "cloudflare_r2", "aws_s3"} {

        if v == "local_fs" && global.Config.LocalFS.Enable {

            up[v] = new(local_fs.LocalFS)

        } else if v == "oss" && global.Config.OSS.Enable {

            c, _ := oss.NewClient()
            up[v] = &oss.OSS{
                Client: c,
            }
        } else if v == "cloudflare_r2" && global.Config.CloudfluR2.Enable {

            c, _ := cloudflare_r2.NewClient()

            up[v] = &cloudflare_r2.R2{
                S3Client: c,
            }
        } else if v == "aws_s3" && global.Config.AWSS3.Enable {

            c, _ := aws_s3.NewClient()

            up[v] = &aws_s3.S3{
                S3Client: c,
            }

        } else {
            continue
        }
        var err error
        dstFileKey, err = up[v].SendFile(fileKey, file, fileHeader)
        if err != nil {
            return nil, err
        }

    }

    accessUrl := global.Config.App.UploadUrlPre + "/" + upload.UrlEscape(dstFileKey)

    return &FileInfo{ImageTitle: fileHeader.Filename, ImageUrl: accessUrl}, nil
}
