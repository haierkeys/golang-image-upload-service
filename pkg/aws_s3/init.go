package aws_s3

import (
    "context"

    "github.com/pkg/errors"

    "github.com/haierspi/golang-image-upload-service/global"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewClient() (*s3.Client, error) {
    // New client

    var region = global.Config.AWSS3.Region
    var accessKeyId = global.Config.AWSS3.AccessKeyID
    var accessKeySecret = global.Config.AWSS3.AccessKeySecret

    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
        config.WithRegion(region),
    )
    if err != nil {
        return nil, errors.Wrap(err, "aws_s3")
    }

    client := s3.NewFromConfig(cfg, func(o *s3.Options) {})

    if err != nil {
        return nil, errors.Wrap(err, "aws_s3")
    }
    return client, nil
}
