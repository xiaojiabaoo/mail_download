package ccl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mail_download/model/request_model"
	"mail_download/service/ccl_ser"
	customErr "mail_download/tools/error"
	"mail_download/tools/response"
	"time"
)

func CCL(ctx *gin.Context) {
	var (
		param  = request_model.CCLParam{}
		err    error
		now    = time.Now()
		serial string
	)
	serial = fmt.Sprintf(`CCL-%d`, now.Unix())
	if err = ctx.ShouldBind(&param); err != nil {
		response.ResponseData(ctx, customErr.New(customErr.PARAMETER_ILLEGAL_DELETION, err.Error()), serial, nil)
		return
	}
	param.Serial = serial
	err = ccl_ser.CLL(param)
	response.ResponseData(ctx, err, param.Serial, param.Serial)
}
