package oss

import (
	"bytes"
	"mime/multipart"
	"strings"

	oss_sdk "github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/haierspi/golang-image-upload-service/global"
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

func (p *OSS) SendFile(fileKey string, f multipart.File, h *multipart.FileHeader) (string, error) {

	if p.Bucket == nil {
		err := p.GetBucket("")
		if err != nil {
			return "", err
		}
	}

	if strings.HasSuffix(global.Config.OSS.CustomPath, "/") {
		fileKey = global.Config.OSS.CustomPath + fileKey
	} else {
		fileKey = global.Config.OSS.CustomPath + "/" + fileKey
	}

	err := p.Bucket.PutObject(fileKey, f)
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
	if strings.HasSuffix(global.Config.OSS.CustomPath, "/") {
		fileKey = global.Config.OSS.CustomPath + fileKey
	} else {
		fileKey = global.Config.OSS.CustomPath + "/" + fileKey
	}

	err := p.Bucket.PutObject(fileKey, bytes.NewReader(content))
	if err != nil {
		return "", err
	}
	return fileKey, nil
}
