package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
func ResponseData(c *gin.Context, err error, data any) {
	if data == "" || data == nil {
		data = make([]int, 0)
	}
	resp := &Response{Data: data, Ret: 0, Msg: custom.ErrText[custom.OPERATION_SUCCESS]}
	if err != nil { // 这里是错误模式
		if tem, ok := err.(*custom.CustomError); ok {
			//var msg = custom.ErrText[tem.Code]
			resp = &Response{Ret: tem.Code, Msg: tem.Error(), Data: data}
		} else {
			// 包装一下
			stack := fmt.Sprintf("stack error trace:\n%+v\n", err) //错误的堆栈
			fmt.Println(stack)
			//返回给前端
			resp = &Response{Ret: custom.SYSTEM_ERROR, Msg: custom.ErrText[custom.SYSTEM_ERROR], Data: make([]int, 0)}
		}
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Current-Time", strconv.Itoa(int(time.Now().Unix())))
		c.JSON(200, resp)
		c.Abort()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Current-Time", strconv.Itoa(int(time.Now().Unix())))
		c.JSON(200, "温馨小提示：具体执行进度可以查看黑色控制板，此页面可以关闭啦！\n1.怎样看程序是否跑完/附件是否下载完：\n   -控制板中提示：所有邮件下载完成，总耗时：x分钟；\n   -只要有这句话显示，就表示程序已经正常执行完毕，所有附件已经下载完成，并且还显示下载附件总共消耗时间\n2.附件下载中途遇到错误会怎样提示：\n   -附件途中遇到错误会跳过，继续下载后面的附件，具体的错误信息会在正常执行完毕后统一提示\n3.附件全部下载完成需要多久：\n   -具体根据邮箱内的所有邮件数量和网络决定，网络越快下载速度也会越快\n   -正常网络下，一个附件下载平均速度为0.5秒\n   -8000个附件的下载时间大约是87分钟")
		c.Abort()
	}

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
