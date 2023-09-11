package code

var (
	//Success                   = NewError(1, "成功")
	Success = NewSuss(1, "成功")

	ErrorNotFoundAPI      = NewError(incr(404), "找不到API")
	ErrorInvalidParams    = NewError(incr(404), "缺少请求参数")
	ErrorTooManyRequests  = NewError(incr(404), "请求过多")
	ErrorInvalidAuthToken = NewError(incr(404), "验证授权Token失败")
	ErrorInvalidToken     = NewError(incr(404), "缺少用户凭证Token")

	ErrorServerInternal = NewError(incr(500), "服务内部错误")
)
