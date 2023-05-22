package download

import (
	"github.com/gin-gonic/gin"
	"mail_download/model/request_model"
	"mail_download/service/download_ser"
	customErr "mail_download/tools/error"
	"mail_download/tools/response"
)

func Download(ctx *gin.Context) {
	var (
		param = request_model.DownloadParam{}
		err   error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		response.ResponseData(ctx, customErr.New(customErr.PARAMETER_ILLEGAL_DELETION, err.Error()), nil)
		return
	}
	err = download_ser.Download(param)
	response.ResponseData(ctx, err, nil)
}
