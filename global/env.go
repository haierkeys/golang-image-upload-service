package global

import (
	"github.com/haierspi/golang-image-upload-service/pkg/path"
)

var (
	// 程序执行目录
	ROOT string
	Name string = "obsidian image-api gateway"
)

func init() {

	filename := path.GetExePath()
	ROOT = filename + "/"

}
