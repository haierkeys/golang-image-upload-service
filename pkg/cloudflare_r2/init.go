package cloudflare_r2

import (
	"context"
	"fmt"
	"log"

	"github.com/haierspi/golang-image-upload-service/global"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewClient() (*s3.Client, error) {
	// New client

	var accountId = global.Config.CloudfluR2.AccountId
	var accessKeyId = global.Config.CloudfluR2.AccessKeyID
	var accessKeySecret = global.Config.CloudfluR2.AccessKeySecret

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
	})

	if err != nil {
		return nil, err
	}
	return client, nil
}
