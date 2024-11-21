package aws_s3

import (
    "bytes"
    "context"
    "fmt"
    "mime/multipart"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/pkg/errors"

    "github.com/haierspi/golang-image-upload-service/global"
    pkg_path "github.com/haierspi/golang-image-upload-service/pkg/path"
)

type S3 struct {
    S3Client  *s3.Client
    S3Manager *manager.Uploader
}

func (p *S3) GetBucket(bucketName string) string {

    // Get bucket
    if len(bucketName) <= 0 {
        bucketName = global.Config.AWSS3.BucketName
    }

    return bucketName
}

// UploadByFile 上传文件
func (p *S3) SendFile(fileKey string, file multipart.File, h *multipart.FileHeader) (string, error) {

    ctx := context.Background()
    bucket := p.GetBucket("")

    fileKey = pkg_path.PathSuffixCheckAdd(global.Config.AWSS3.CustomPath, "/") + fileKey

    k, _ := h.Open()

    _, err := p.S3Client.PutObject(ctx, &s3.PutObjectInput{
        Bucket:      aws.String(bucket),
        Key:         aws.String(fileKey),
        Body:        k,
        ContentType: aws.String(h.Header.Get("Content-Type")),
    })

    if err != nil {
        return "", errors.Wrap(err, "aws_s3")
    }

    return fileKey, nil
}

func (p *S3) SendContent(fileKey string, content []byte) (string, error) {

    ctx := context.Background()
    bucket := p.GetBucket("")

    fileKey = pkg_path.PathSuffixCheckAdd(global.Config.AWSS3.CustomPath, "/") + fileKey

    input := &s3.PutObjectInput{
        Bucket:            aws.String(bucket),
        Key:               aws.String(fileKey),
        Body:              bytes.NewReader(content),
        ChecksumAlgorithm: types.ChecksumAlgorithmSha256,
    }
    output, err := p.S3Manager.Upload(ctx, input)
    if err != nil {
        var noBucket *types.NoSuchBucket
        if errors.As(err, &noBucket) {
            fmt.Printf("Bucket %s does not exist.\n", bucket)
            err = noBucket
        }
    } else {
        err := s3.NewObjectExistsWaiter(p.S3Client).Wait(ctx, &s3.HeadObjectInput{
            Bucket: aws.String(bucket),
            Key:    aws.String(fileKey),
        }, time.Minute)
        if err != nil {
            fmt.Printf("Failed attempt to wait for object %s to exist in %s.\n", fileKey, bucket)
        } else {
            _ = *output.Key
        }
    }

    return fileKey, errors.Wrap(err, "aws_s3")
}
