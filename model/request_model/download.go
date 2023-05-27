package request_model

type DownloadParam struct {
	Server   string `form:"server"`   // 邮箱服务器地址
	Account  string `form:"account"`  // 邮箱账号
	Password string `form:"password"` // 邮箱密码
	Time     string `form:"time"`     // 下载邮件的时间范围，比如：2023-01-01 00:00:00，表示下载这个时间（>=这个时间）到当前时间的所有邮件附件，该字段和count字段只能二选一
	Count    uint32 `form:"count"`    // 下载指定数量的邮件，比如：10表示下载最新的十封邮件；20表示下载最新的二十封邮件；适合第一次下载附件使用；该字段和time字段只能二选一
	Url      string `form:"url"`      // 附件下载的所在目录；如果目录不存在，会在指定位置创建该目录
	Type     string `form:"type"`     // 文件操作类型，附件pdf文件存在时的操作逻辑：cover.覆盖；jump.跳过；all_reserved.全保留
}
