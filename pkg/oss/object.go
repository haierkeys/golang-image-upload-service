package oss

import (
	"github.com/haierspi/golang-image-upload-service/global"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func newClient() (*oss.Client, error) {
	// New client
	client, err := oss.New(global.OSSSetting.Endpoint, global.OSSSetting.AccessKeyID, global.OSSSetting.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetBucket(bucketName string) (*oss.Bucket, error) {

	client, _ := newClient()
	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
