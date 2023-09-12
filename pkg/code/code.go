package code

import (
	"fmt"
	"net/http"
)

type Code struct {
	// 状态码
	code int
	// 状态
	status bool
	// 错误消息
	msg string
	// 数据
	data interface{}
	// 错误详细信息
	details []string
	// 是否含有详情
	haveDetails bool
}

var codes = map[int]string{}
var maxcode = 0

func NewError(code int, msg string) *Code {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}

	codes[code] = msg

	if code > maxcode {
		maxcode = code
	}

	return &Code{code: code, status: false, msg: msg}
}

func incr(code int) int {
	if code > maxcode {
		return code
	} else {
		return maxcode + 1
	}
}

var sussCodes = map[string]string{}

func NewSuss(code int, msg string) *Code {
	if _, ok := sussCodes[msg]; ok {
		panic(fmt.Sprintf("成功信息 %d 已经存在，请更换一个", msg))
	}
	sussCodes[msg] = msg
	if code > maxcode {
		maxcode = code
	}
	return &Code{code: code, status: true, msg: msg}
}

func (e *Code) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e *Code) Code() int {
	return e.code
}

func (e *Code) Status() bool {
	return e.status
}

func (e *Code) Msg() string {
	return e.msg
}

func (e *Code) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Code) Details() []string {
	return e.details
}

func (e *Code) Data() interface{} {
	return e.data
}

func (e *Code) HaveDetails() bool {
	return e.haveDetails
}

func (e *Code) WithData(data interface{}) *Code {
	e.data = data
	return e
}

func (e *Code) WithDetails(details ...string) *Code {
	e.haveDetails = true
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}

	return e
}

func (e *Code) StatusCode() int {
	return http.StatusOK
}
