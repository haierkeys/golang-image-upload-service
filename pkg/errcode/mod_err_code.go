package errcode

var (
	ErrorUploadFileFail = NewError(20001, "上传文件失败")

	ErrorOrderViewFail = NewError(23021, "订单无法查看失败")
)
