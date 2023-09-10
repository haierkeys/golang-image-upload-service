package service

import (
	"github.com/haierspi/golang-image-upload-service/pkg/timef"
)

type ClientVersion struct {
	Id          int64      `json:"ID"`          // 用户id
	Platform    int32      `json:"platform"`    // 平台 1:安卓 2:ios
	VersionCode int64      `json:"versionCode"` // 版本号
	VersionName string     `json:"versionName"` // 版本名
	Details     string     `json:"details"`     // 版本更新详情
	ResourceUrl string     `json:"resourceURL"` // 版本更新资源地址
	CreatedAt   timef.Time `json:"createdAt"`   //创建时间
	UpdatedAt   timef.Time `json:"updatedAt"`   //更新时间
}

// ClientVersionRequest 客户端查询更新请求参数
type ClientVersionRequest struct {
	Platform    int32 `form:"platform" binding:"required,gte=0"`    // 平台 1:安卓 2:ios
	VersionCode int64 `form:"versionCode" binding:"required,gte=0"` // 版本号
}

// ClientVersion 客户端查询更新
func (svc *Service) ClientVersion(param *ClientVersionRequest) (*ClientVersion, error) {

	return &ClientVersion{}, nil

}
