package oss

import (
    "bytes"
    "io"

    oss_sdk "github.com/aliyun/aliyun-oss-go-sdk/oss"

    "github.com/haierspi/golang-image-upload-service/global"
    pkg_path "github.com/haierspi/golang-image-upload-service/pkg/path"
)

type OSS struct {
    Client *oss_sdk.Client
    Bucket *oss_sdk.Bucket
}

func (p *OSS) GetBucket(bucketName string) error {

    // Get bucket
    if len(bucketName) <= 0 {
        bucketName = global.Config.OSS.BucketName
    }
    var err error
    p.Bucket, err = p.Client.Bucket(bucketName)
    return err
}

func (p *OSS) SendFile(fileKey string, file io.Reader, itype string) (string, error) {

    if p.Bucket == nil {
        err := p.GetBucket("")
        if err != nil {
            return "", err
        }
    }

    fileKey = pkg_path.PathSuffixCheckAdd(global.Config.OSS.CustomPath, "/") + fileKey

    err := p.Bucket.PutObject(fileKey, file)
    if err != nil {
        return "", err
    }
    return fileKey, nil
}

func (p *OSS) SendContent(fileKey string, content []byte) (string, error) {

    if p.Bucket == nil {
        err := p.GetBucket("")
        if err != nil {
            return "", err
        }
    }

    fileKey = pkg_path.PathSuffixCheckAdd(global.Config.OSS.CustomPath, "/") + fileKey

    err := p.Bucket.PutObject(fileKey, bytes.NewReader(content))
    if err != nil {
        return "", err
    }
    return fileKey, nil
}
