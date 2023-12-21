package download

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mail_download/model/request_model"
	"mail_download/service/download_ser"
	customErr "mail_download/tools/error"
	"mail_download/tools/response"
	"time"
)

func Download(ctx *gin.Context) {
	var (
		param  = request_model.DownloadParam{}
		err    error
		now    = time.Now()
		serial string
	)
	serial = fmt.Sprintf(`MAIL-%d`, now.Unix())
	if err = ctx.ShouldBind(&param); err != nil {
		response.ResponseData(ctx, customErr.New(customErr.PARAMETER_ILLEGAL_DELETION, err.Error()), serial, nil)
		return
	}
	param.Serial = serial
	err = download_ser.Download(param)
	response.ResponseData(ctx, err, serial, serial)
}
