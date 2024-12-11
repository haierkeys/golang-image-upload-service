package service

import (
    "bytes"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "io"
    "mime/multipart"

    "github.com/disintegration/imaging"
    "github.com/gen2brain/avif"
    "github.com/google/uuid"
    _ "github.com/gookit/goutil/dump"
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

type ClientUploadParams struct {
    Key    string `form:"key"`
    Type   string `form:"type"`
    Width  int    `form:"width"`
    Height int    `form:"height"`
}

type Uploader interface {
    SendFile(pathKey string, file io.Reader, cType string) (string, error)
    SendContent(pathKey string, content []byte) (string, error)
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader, form *ClientUploadParams) (*FileInfo, error) {

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

    // 默认裁剪 | 居中裁剪 | 固定尺寸拉伸 | 固定尺寸等比缩放不裁切 | 不处理
    // type: "fill-topleft" | "fill-center" | "resize" | "fit" | "none";

    // 服务器强制限制图片的宽度和高度
    var imageMaxWidth = global.Config.App.ImageMaxSizeWidth
    var imageMaxHeight = global.Config.App.ImageMaxSizeHeight
    var newWidth, newHeight int
    var newImage image.Image
    var isNewImage bool

    if form.Type == "none" || form.Type == "" {

        newWidth = imageMaxWidth
        newHeight = imageMaxHeight

        if (size.X != newWidth || size.Y != newHeight) && (newWidth != 0 || newHeight != 0) {

            if newWidth == 0 || newHeight == 0 {
                newImage = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
            } else {
                newImage = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)
            }

            isNewImage = true
        }
    } else if form.Type == "fill-topleft" {
        if form.Width < imageMaxWidth || imageMaxWidth == 0 {
            newWidth = form.Width
        } else {
            newWidth = imageMaxWidth
        }
        if form.Height < imageMaxHeight || imageMaxHeight == 0 {
            newHeight = form.Height
        } else {
            newHeight = imageMaxHeight
        }

        newImage = imaging.Fill(img, newWidth, newHeight, imaging.TopLeft, imaging.Lanczos)
        isNewImage = true
    } else if form.Type == "fill-center" {
        if form.Width < imageMaxWidth || imageMaxWidth == 0 {
            newWidth = form.Width
        } else {
            newWidth = imageMaxWidth
        }
        if form.Height < imageMaxHeight || imageMaxHeight == 0 {
            newHeight = form.Height
        } else {
            newHeight = imageMaxHeight
        }
        // newImage = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)
        newImage = imaging.Fill(img, newWidth, newHeight, imaging.Center, imaging.Lanczos)
        isNewImage = true
    } else if form.Type == "resize" {

        if form.Width < imageMaxWidth || imageMaxWidth == 0 {
            newWidth = form.Width
        } else {
            newWidth = imageMaxWidth
        }
        if form.Height < imageMaxHeight || imageMaxHeight == 0 {
            newHeight = form.Height
        } else {
            newHeight = imageMaxHeight
        }

        if form.Width != 0 && form.Height != 0 && (size.X != newWidth || size.Y != newHeight) {
            newImage = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
            isNewImage = true
        }
    } else if form.Type == "fit" {

        if form.Width < imageMaxWidth || imageMaxWidth == 0 {
            newWidth = form.Width
        } else {
            newWidth = imageMaxWidth
        }
        if form.Height < imageMaxHeight || imageMaxHeight == 0 {
            newHeight = form.Height
        } else {
            newHeight = imageMaxHeight
        }

        if (size.X != newWidth || size.Y != newHeight) && (newWidth != 0 || newHeight != 0) {

            if newWidth == 0 || newHeight == 0 {
                newImage = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
            } else {
                newImage = imaging.Fit(img, newWidth, newHeight, imaging.Lanczos)
            }

            isNewImage = true
        }
    }

    if isNewImage {

        // 调整图片大小

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
        file.Seek(0, 0)
        io.Copy(writer, file)
    }

    reader := bytes.NewReader(writer.Bytes())

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

        reader.Seek(0, 0)

        var err error
        dstFileKey, err = up[v].SendFile(fileKey, reader, cType)
        if err != nil {
            return nil, err
        }

    }

    accessUrl := pkg_path.PathSuffixCheckAdd(global.Config.App.UploadUrlPre, "/") + upload.UrlEscape(dstFileKey)

    return &FileInfo{ImageTitle: fileHeader.Filename, ImageUrl: accessUrl}, nil
}
func MemDupReader(r io.Reader) func() io.Reader {
    b := bytes.NewBuffer(nil)
    t := io.TeeReader(r, b)

    return func() io.Reader {
        br := bytes.NewReader(b.Bytes())
        return io.MultiReader(br, t)
    }
}
