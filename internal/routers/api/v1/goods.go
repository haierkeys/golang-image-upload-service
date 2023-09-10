package v1

import (
	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/internal/service"
	"github.com/haierspi/golang-image-upload-service/pkg/app"
	"github.com/haierspi/golang-image-upload-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Goods struct{}

func NewGoods() Goods {
	return Goods{}
}

func (t *Goods) List(c *gin.Context) {

	param := &service.GoodsListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	svc := service.New(c.Request.Context())
	svcList, svcCount, err := svc.GoodsList(param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponseList(svcList, svcCount)
	return
}

func (t *Goods) Details(c *gin.Context) {

	param := &service.GoodsDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcData, err := svc.GoodsDetails(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponse(svcData)
	return
}
