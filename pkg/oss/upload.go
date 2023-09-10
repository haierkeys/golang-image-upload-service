package oss

import (
	"io"
	"net/http"
	"strings"

	"github.com/haierspi/golang-image-upload-service/global"
)

func UploadByString(path string, content string) error {

	bucket, err := GetBucket(global.OSSSetting.BucketName)

	if err != nil {
		return err
	}

	err = bucket.PutObject(path, strings.NewReader(content))
	if err != nil {
		return err
	}
	return nil
}

func UploadByURL(path string, url string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		return err2
	}

	bucket, err := GetBucket(global.OSSSetting.BucketName)

	if err != nil {
		return err
	}

	err = bucket.PutObject(path, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	return nil
}
