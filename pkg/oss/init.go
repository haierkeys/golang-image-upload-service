package oss

import (
	"github.com/haierspi/golang-image-upload-service/global"

	oss_sdk "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client map[string]*oss_sdk.Client

func NewClient() (*oss_sdk.Client, error) {

	id := global.Config.OSS.AccessKeyID
	var err error
	if client[id] != nil {
		return client[id], nil
	}
	// New client
	client[id], err = oss_sdk.New(global.Config.OSS.Endpoint, global.Config.OSS.AccessKeyID, global.Config.OSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return client[id], nil
}
