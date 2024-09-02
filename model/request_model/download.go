package request_model

type DownloadParam struct {
	Server        string   `json:"server"`         // 邮箱服务器地址，此参数必须填写
	Account       string   `json:"account"`        // 邮箱账号，此参数必须填写
	Password      string   `json:"password"`       // 邮箱密码，此参数必须填写
	Box           []string `json:"box"`            // 收件箱
	TimeLimit     int64    `json:"time_limit"`     // 是否下载指定时间范围的邮件：1.是；0.否；
	Time          int64    `json:"time"`           // 下载邮件的时间范围，表示下载这个时间（>=这个时间）到当前时间的所有邮件附件，该字段和count字段只能二选一
	EndTime       int64    `json:"end_time"`       // 结束时间，如果 TimeLimit = 0，该值无效
	Count         uint32   `json:"count"`          // 下载指定数量的邮件，比如：10表示下载最新的十封邮件；20表示下载最新的二十封邮件；适合第一次下载附件使用；该字段和time字段只能二选一
	Url           string   `json:"url"`            // 附件下载的所在目录；如果目录不存在，会在指定位置创建该目录
	Classify      string   `json:"classify"`       // 下载目录划分级别：year.年；month.月；day.天；如果选择的是year，那么邮件附件是按年份的目录存放；如果选择的是month那么邮件附件是按年份+月份的目录存放；如果选择的是day那么邮件附件是按年份+月份+天的目录存放，此参数必须填写
	Type          string   `json:"type"`           // 文件操作类型，附件pdf文件存在时的操作逻辑：cover.覆盖；jump.跳过；all_reserved.全保留；注意：如果是classify为year，表示操作的是在一年内的同名文件；如果是classify为month，表示操作的是在一个月内的同名文件；如果是classify为day以此类推，此参数必须填写
	Inform        string   `json:"inform"`         // 是否需要通知：on.通知；off.不通知
	InformAccount string   `json:"inform_account"` // 通知账号，如果为空默认使用登录的邮箱来通知
	Serial        string   `json:"-"`              // 操作流水号，后台使用
	Downloaded    string   `json:"downloaded"`     // 是否需要将邮件添加到 “已下载” 文件夹中：on.添加；off.不添加
}

type MailboxesParam struct {
	Server   string `form:"server" binding:"required"`   // 邮箱服务器地址，此参数必须填写
	Account  string `form:"account" binding:"required"`  // 邮箱账号，此参数必须填写
	Password string `form:"password" binding:"required"` // 邮箱密码，此参数必须填写
}

type CCLParam struct {
	Account  string `json:"account" binding:"required"`  // 登录账号，此参数必须填写
	Password string `json:"password" binding:"required"` // 登录密码，此参数必须填写
	Url      string `json:"url" binding:"required"`      // 附件下载的所在目录；如果目录不存在，会在指定位置创建该目录
	Email    string `json:"email"`                       // 通知邮箱账号
	Serial   string `json:"-"`                           // 操作流水号，后台使用
}
