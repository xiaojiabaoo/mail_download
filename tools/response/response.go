package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mail_download/tools"
	custom "mail_download/tools/error"
	"strconv"
	"time"
)

type Response struct {
	Ret  int64  `json:"ret"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 返回响应
func ResponseData(c *gin.Context, err error, serial string, data any) {
	if data == "" || data == nil {
		data = make([]int, 0)
	}
	resp := &Response{Data: data, Ret: 0, Msg: custom.ErrText[custom.OPERATION_SUCCESS]}
	if err != nil { // 这里是错误模式
		if tem, ok := err.(*custom.CustomError); ok {
			//var msg = custom.ErrText[tem.Code]
			go tools.Logger(serial, tem.Error(), tools.LOG_LEVEL_ERROR)
			resp = &Response{Ret: tem.Code, Msg: tem.Error(), Data: data}
		} else {
			// 包装一下
			stack := fmt.Sprintf("stack error trace:\n%+v\n", err) //错误的堆栈
			fmt.Println(stack)
			go tools.Logger(serial, err.Error(), tools.LOG_LEVEL_SYSTEM_ERROR)
			//返回给前端
			resp = &Response{Ret: custom.SYSTEM_ERROR, Msg: custom.ErrText[custom.SYSTEM_ERROR], Data: make([]int, 0)}
		}
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Current-Time", strconv.Itoa(int(time.Now().Unix())))
	c.JSON(200, resp)
	c.Abort()
}

func ResponseInteger(c *gin.Context, data interface{}) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Current-Time", strconv.Itoa(int(time.Now().Unix())))
	c.Header("application/json;charset=UTF-8", strconv.Itoa(int(time.Now().Unix())))
	c.JSON(200, data)
	c.Abort()
}

func ResponseXml(c *gin.Context, data []byte) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Current-Time", strconv.Itoa(int(time.Now().Unix())))
	c.Header("application/json;charset=UTF-8", strconv.Itoa(int(time.Now().Unix())))
	c.Writer.Write(data)
	c.Abort()
}
