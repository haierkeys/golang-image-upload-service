package app

import (
	"net/http"
	"reflect"

	"github.com/haierspi/golang-image-upload-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"pageSize"`
	// 总行数
	TotalRows int `json:"totalRows"`
}
type ListRes struct {
	// 数据清单
	List interface{} `json:"list"`
	// 翻页信息
	Pager Pager `json:"pager"`
}

type ResResult struct {
	// HTTP状态码
	Code interface{} `json:"code"`
	// 业务状态码
	Status interface{} `json:"status"`
	// 失败&&成功消息
	Msg interface{} `json:"msg"`
	// 数据集合
	Data interface{} `json:"data"`
}

type ResListResult struct {
	// HTTP状态码
	Code interface{} `json:"code"`
	// 业务状态码
	Status interface{} `json:"status"`
	// 失败&&成功消息
	Msg interface{} `json:"msg"`
	// 数据集合
	Data ListRes `json:"data"`
}

type ErrResult struct {
	// HTTP状态码
	Code interface{} `json:"code"`
	// 业务状态码
	Status interface{} `json:"status"`
	// 失败&&成功消息
	Msg interface{} `json:"msg"`
	// 错误格式数据
	Data interface{} `json:"data"`
	// 错误支付
	Details interface{} `json:"details"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

// 解析gin中参数 并设置为字符串属性 用于判断参数类型零值的问题
//
//	type ListOrderRequest struct {
//		Status    int64  `form:"status" request:"StatusStr"` //状态
//		StatusStr string //会自动根据上面的 request tag 设置这个值
//	}
func RequestParamStrParse(c *gin.Context, param any) {
	tParam := reflect.TypeOf(param).Elem()
	vParam := reflect.ValueOf(param).Elem()
	for i := 0; i < tParam.NumField(); i++ {
		name := tParam.Field(i).Name
		if nameType, ok := tParam.FieldByName(name); ok {
			dstName := nameType.Tag.Get("request")
			if dstName != "" {
				paramName := nameType.Tag.Get("form")
				if value, ok := c.GetQuery(paramName); ok {
					vParam.FieldByName(dstName).SetString(value)
				}
			}
		}
	}
}

// 获取ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

func GetAccessHost(c *gin.Context) string {
	AccessProto := ""
	if proto := c.Request.Header.Get("X-Forwarded-Proto"); proto == "" {
		AccessProto = "http" + "://"
	} else {
		AccessProto = proto + "://"
	}
	return AccessProto + c.Request.Host
}

// 输出到浏览器
func (r *Response) ToResponse(data interface{}) {
	code := errcode.Success
	r.SendResponse(http.StatusOK, ResResult{
		Code:   http.StatusOK,
		Status: code.Code(),
		Msg:    code.Msg(),
		Data:   data,
	})

}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	code := errcode.Success
	r.SendResponse(http.StatusOK, ResListResult{
		Code:   http.StatusOK,
		Status: code.Code(),
		Msg:    code.Msg(),
		Data: ListRes{
			List: list,
			Pager: Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	r.SendResponse(err.StatusCode(), ErrResult{
		Code:    http.StatusOK,
		Status:  err.Code(),
		Msg:     err.Msg(),
		Data:    err.Data(),
		Details: err.Details(),
	})
}

func (r *Response) ToErrorResponseData(err *errcode.Error, errSrc error) {
	r.SendResponse(err.StatusCode(), ErrResult{
		Code:   http.StatusOK,
		Status: err.Code(),
		Msg:    err.Msg(),
		Data:   errSrc.Error(),
	})
}

func (r *Response) SendResponse(statusCode int, content interface{}) {

	var debug string
	var exist bool

	debug, exist = r.Ctx.GetQuery("debug")

	if !exist {
		debug = r.Ctx.GetHeader("debug")
	}

	if debug != "" {

	}

	r.Ctx.JSON(http.StatusOK, content)

}
