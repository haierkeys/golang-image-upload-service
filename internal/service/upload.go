package service

import (
    "bytes"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "io"
    "mime/multipart"

    "github.com/gen2brain/avif"
    "github.com/google/uuid"
    "github.com/nfnt/resize"
    "github.com/pkg/errors"
    "golang.org/x/image/bmp"
    "golang.org/x/image/tiff"
    _ "golang.org/x/image/webp"

    "github.com/haierspi/golang-image-upload-service/global"
    "github.com/haierspi/golang-image-upload-service/pkg/aws_s3"
    "github.com/haierspi/golang-image-upload-service/pkg/cloudflare_r2"
    "github.com/haierspi/golang-image-upload-service/pkg/local_fs"
    "github.com/haierspi/golang-image-upload-service/pkg/oss"
    pkg_path "github.com/haierspi/golang-image-upload-service/pkg/path"
    "github.com/haierspi/golang-image-upload-service/pkg/upload"
)

type FileInfo struct {
    ImageTitle string `json:"imageTitle"`
    ImageUrl   string `json:"imageUrl"`
}

type Uploader interface {
    SendFile(pathKey string, file io.Reader, cType string) (string, error)
    SendContent(pathKey string, content []byte) (string, error)
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
    return svc.fileSyncHandle(fileType, file, fileHeader)
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

    cType := fileHeader.Header.Get("Content-Type")

    if !upload.CheckContainExt(fileType, fileName) {
        return nil, errors.New("file suffix is not supported.")
    }
    if upload.CheckMaxSize(fileType, file) {
        return nil, errors.New("exceeded maximum file limit.")
    }

    fileKey := upload.GetSavePreDirPath() + fileName

    var up = make(map[string]Uploader)
    var dstFileKey string

    writer := &bytes.Buffer{}

    // 压缩
    _, err := file.Seek(0, 0)

    img, filetype, err := image.Decode(file)

    if err != nil {
        return nil, err
    }

    size := img.Bounds().Size()

    if size.X > global.Config.App.ImageMaxSizeWidth || size.Y > global.Config.App.ImageMaxSizeHeight {

        wRatio := float64(size.X) / float64(global.Config.App.ImageMaxSizeWidth)
        hRatio := float64(size.Y) / float64(global.Config.App.ImageMaxSizeHeight)
        ratio := wRatio
        if hRatio > wRatio {
            ratio = hRatio
        }
        // 创建新的图像对象
        newWidth := uint(float64(size.X) / ratio)
        newHeight := uint(float64(size.Y) / ratio)

        // 调整图片大小
        newImage := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

        switch filetype {
        case "png":
            err = png.Encode(writer, newImage)
        case "gif":
            err = gif.Encode(writer, newImage, &gif.Options{NumColors: 256})
        case "jpeg", "jpg":
            err = jpeg.Encode(writer, newImage, &jpeg.Options{Quality: global.Config.App.ImageQuality})
        case "bmp":
            err = bmp.Encode(writer, newImage)
        case "tif", "tiff":
            err = tiff.Encode(writer, newImage, nil)
        case "webp":
            cType = "image/jpg"
            ext := upload.GetFileExt(fileKey)
            fileKey = fileKey[0:len(fileKey)-len(ext)] + ".jpg"

            err = jpeg.Encode(writer, newImage, &jpeg.Options{Quality: global.Config.App.ImageQuality})
        case "avif":
            err = avif.Encode(writer, newImage, avif.Options{Quality: global.Config.App.ImageQuality})

        default:
            return nil, errors.New("Unknown image type:" + filetype)
        }

        if err != nil {
            return nil, err
        }

    } else {
        _, err = io.Copy(writer, file)
    }

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
        dstFileKey, err = up[v].SendFile(fileKey, writer, cType)
        if err != nil {
            return nil, err
        }

    }

    accessUrl := pkg_path.PathSuffixCheckAdd(global.Config.App.UploadUrlPre, "/") + upload.UrlEscape(dstFileKey)

    return &FileInfo{ImageTitle: fileHeader.Filename, ImageUrl: accessUrl}, nil
}
