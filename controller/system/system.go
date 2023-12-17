package system

import (
	"github.com/gin-gonic/gin"
	"mail_download/model/request_model"
	"mail_download/service/system_ser"
	customErr "mail_download/tools/error"
	"mail_download/tools/response"
	"time"
)

func List(ctx *gin.Context) {
	var (
		param = request_model.SysListParam{}
		err   error
		now   = time.Now()
	)
	if err = ctx.ShouldBind(&param); err != nil {
		response.ResponseData(ctx, customErr.New(customErr.PARAMETER_ILLEGAL_DELETION, err.Error()), now.Unix(), nil)
		return
	}
	list, err := system_ser.List(param)
	response.ResponseData(ctx, err, now.Unix(), list)
}
