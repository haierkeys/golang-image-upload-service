package service

import (
	"github.com/haierspi/golang-image-upload-service/pkg/app"
	"github.com/haierspi/golang-image-upload-service/pkg/timef"
)

type Goods struct {
	GoodsId            int64      `json:"goodsID"`            // 商品ID
	GoodsSn            string     `json:"goodsSN"`            // 货号
	CategoryId         int64      `json:"categoryID"`         // 一级分类
	CategoryName       string     `json:"categoryName"`       // 一级分类
	CopyrightId        int64      `json:"copyrightID"`        // 版权方ID
	CopyrightName      string     `json:"copyrightName"`      // 版权方名字
	CopyrightImage     string     `json:"copyrightImage"`     // 版权方图片
	BrandId            int64      `json:"brandID"`            // 品牌ID
	BrandName          string     `json:"brandName"`          // 品牌方名称
	ReleaseId          int64      `json:"releaseID"`          // 发行方ID
	ReleaseName        string     `json:"releaseName"`        // 发行方名称
	GoodsType          int32      `json:"goodsType"`          // 0：实物商品，1：数字藏品，2：实物礼包， 3：数字藏品-盲盒
	GoodsName          string     `json:"goodsName"`          // 商品名称
	GoodsPrice         float64    `json:"goodsPrice"`         // 商品价格
	GoodsMarketPrice   float64    `json:"goodsMarketPrice"`   // 市场价
	GoodsExpressType   int32      `json:"goodsExpressType"`   // 运费类型 0:免运费，1:全国运费一个价, 2:根据地区和货物重量单独计算）
	GoodsExpressPrice  float64    `json:"goodsExpressPrice"`  // 商品运费 0 免运费
	BuyNumLimit        int64      `json:"buyNumLimit"`        // 限购数量
	GoodsUrl           string     `json:"goodsURL"`           // 连接关键词(英文)
	GoodsWeight        int64      `json:"goodsWeight"`        // 商品单位重量（g）
	GoodsStock         int64      `json:"goodsStock"`         // 当前库存
	GoodsTotalStock    int64      `json:"goodsTotalStock"`    // 历史总库存
	GoodsTitlePic      string     `json:"goodsTitlePic"`      // 商品标题图
	GoodsThumbPic      string     `json:"goodsThumbPic"`      // 商品缩略图
	GoodsImage         string     `json:"goodsImage"`         // 商品封面图
	GoodsAr            string     `json:"goodsAr"`            // ar 模型
	GoodsArImage       string     `json:"goodsArImage"`       // ar加载图
	GoodsTags          string     `json:"goodsTags"`          // 标签 使用英文逗号间隔
	GoodsBody          string     `json:"goodsBody"`          // 商品内容
	GoodsBodyMobile    string     `json:"goodsBodyMobile"`    // 商品内容移动版
	ContractTemplateId int64      `json:"contractTemplateID"` // 合约类型模板ID
	BlockchainId       int64      `json:"blockchainID"`       // 区块链类型
	BlockchainName     string     `json:"blockchainName"`     // 区块链名字
	BlockchainKey      string     `json:"blockchainKey"`      // 区块链key
	BlockchainIcon     string     `json:"blockchainIcon"`     // 区块链ICON
	BlockchainAddress  string     `json:"blockchainAddress"`  // 区块链地址
	Status             int32      `json:"status"`             // 商品状态（0下架,1上架）
	Weight             int64      `json:"weight"`             // 排序权重
	ReleasedTime       int64      `json:"releasedTime"`       // 发行时间(时间戳)
	SaleTime           int64      `json:"saleTime"`           // 销售时间(时间戳)
	ReleasedAt         timef.Time `json:"releasedAt"`         // 发行时间
	SaleAt             timef.Time `json:"saleAt"`             // 销售时间
	CreatedAt          timef.Time `json:"createdAt"`          // 创建时间
	UpdatedAt          timef.Time `json:"updatedAt"`          // 更新时间
}

// GoodsDetailsRequest 根据url商品详情请求参数
type GoodsDetailsRequest struct {
	GoodsID int64 `form:"goodsID" binding:"required"` //商品ID
}

// GoodsDetailsURLRequest 根据url商品详情请求参数
type GoodsDetailsURLRequest struct {
	GoodsURL string `form:"goodsURL" binding:"required"`
}

// GoodsListRequest 商品列表请求参数
type GoodsListRequest struct {
	CategoryID int64 `form:"categoryID"` //分类信息
}

// GoodsList 获取商品列表
func (svc *Service) GoodsList(param *GoodsListRequest, pager *app.Pager) ([]*Goods, int, error) {

	return nil, 0, nil
}

// GoodsDetails 根据商品ID获取商品详情
func (svc *Service) GoodsDetails(param *GoodsDetailsRequest) (*Goods, error) {

	return &Goods{}, nil
}
