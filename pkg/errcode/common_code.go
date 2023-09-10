package errcode

var (
	//Success                   = NewError(1, "成功")
	Success = NewSuss(1, "成功")

	ServerError               = NewError(10001, "服务内部错误")
	InvalidParams             = NewError(10002, "缺少请求参数")
	NotFound                  = NewError(10003, "找不到")
	UnauthorizedAuthNotExist  = NewError(10004, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10005, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10006, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10007, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10008, "请求过多")
	InvalidToken              = NewError(10009, "缺少用户凭证Token")
)
