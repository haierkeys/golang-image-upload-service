package global

import (
	"github.com/haierspi/golang-image-upload-service/pkg/path"
)

var (
	//程序执行目录
	ROOT   string
	CONFIG string
)

const (
	MqOrderNftMitKey            string = "order-nft-mit"                                                    // NFT Mint
	MqOrderNftMetadataChangeKey string = "order-nft-metadata-change"                                        // NFT Metadata Change
	MqOrderNftAppKey            string = "order-nft-app"                                                    // APP通知操作
	MqOrderNftGivenKey          string = "order-nft-given"                                                  // Nft转赠
	NftAccountAddr              string = ""                   // NFT操作账户
	NftAccountPrivateKey        string = "" // NFT操作账户私钥
)

func init() {

	filename := path.GetExePath()
	ROOT = filename + "/"

}
