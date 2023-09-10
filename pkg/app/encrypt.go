package app

import (
	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/pkg/util"
)

func Convert(str string) string {
	key := global.SecuritySetting.HtmlEncryptKey

	tokenByte := []rune(str)
	keyByte := []rune(key)
	tmpCode := util.XorEncodeStrRune(tokenByte, keyByte)

	return string(tmpCode)
}
