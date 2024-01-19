package system

import (
	"github.com/gin-gonic/gin"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/service/system_ser"
	customErr "mail_download/tools/error"
	"mail_download/tools/response"
)

func List(ctx *gin.Context) {
	var (
		param = request_model.SysListParam{}
		err   error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		response.ResponseData(ctx, customErr.New(customErr.PARAMETER_ILLEGAL_DELETION, err.Error()), "", nil)
		return
	}
	list, err := system_ser.List(param)
	response.ResponseData(ctx, err, "", list)
}

func Update(ctx *gin.Context) {
	var (
		err error
	)
	err = system_ser.Update()
	response.ResponseData(ctx, err, "", nil)
}

func CheckUpdate(ctx *gin.Context) {
	var (
		param  = request_model.CheckUpdateParam{}
		err    error
		update response_model.Version
	)
	if err = ctx.ShouldBind(&param); err != nil {
		response.ResponseData(ctx, customErr.New(customErr.PARAMETER_ILLEGAL_DELETION, err.Error()), "", nil)
		return
	}
	update, err = system_ser.CheckUpdate(param)
	response.ResponseData(ctx, err, "", update)
}
