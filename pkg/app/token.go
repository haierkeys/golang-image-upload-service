package app

import (
	"encoding/json"

	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/pkg/util"

	"github.com/gin-gonic/gin"
)

type UserEntity struct {
	Uid    int64  `json:"uid"`
	Expiry int64  `json:"expiry"`
	IP     string `json:"ip"`
}

func ParseToken(token string) (*UserEntity, error) {

	var userEntity UserEntity

	resultStr, err := util.AuthDzCodeEncrypt(token, "DECODE", global.SecuritySetting.TokenAuthKey, 0)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resultStr), &userEntity)

	if err == nil {
		return &userEntity, nil
	}

	return nil, err

}

func GetUid(ctx *gin.Context) (out int64) {
	user, exist := ctx.Get("user_token")

	if exist == true {
		out = user.(*UserEntity).Uid
	}
	return
}

func GetIP(ctx *gin.Context) (out string) {
	user, exist := ctx.Get("user_token")
	if exist == true {
		out = user.(*UserEntity).IP
	}
	return out
}

func GetExpiration(ctx *gin.Context) (out int64) {
	user, exist := ctx.Get("user_token")
	if exist == true {
		out = user.(*UserEntity).Expiry
	}
	return out
}
