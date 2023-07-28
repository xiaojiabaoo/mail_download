package error

type CustomError struct {
	err  string
	Code int64
}

func (m *CustomError) Error() string {
	return m.err
}

func New(code int64, msg string) error {
	if msg == "" {
		msg = GetErrMsg(code)
	}
	return &CustomError{
		err:  msg,
		Code: code,
	}
}

const (
	OPERATION_SUCCESS          = 0
	SYSTEM_ERROR               = 10001
	TIME_OUT                   = 10002
	FUNCTION_NOT_AVAILABLE     = 10003
	PARAMETER_ILLEGAL_DELETION = 10004

	DATA_NOT_EXIST = 20001

	SERVER_ADDR_ERROR            = 30001
	EMAIL_ACCOUNT_ERROR          = 30002
	EMAIL_PASSWORD_ERROR         = 30003
	TIME_COUNT_ERROR             = 30004
	EMAIL_DOWNLOAD_URL_ERROR     = 30005
	EMAIL_LOGIN_ERROR            = 30006
	DOWNLOAD_URL_ERROR           = 30007
	DOWNLOAD_TIME_ERROR          = 30008
	EMAIL_EMPTY_ERROR            = 30009
	FILE_TYPE_EMPTY_ERROR        = 30010
	FILE_DOWNLOAD_CLASSIFY_ERROR = 30011
	GET_SYSTEM_PATH_ERROR        = 30012
)

var ErrText = map[int64]string{
	OPERATION_SUCCESS:          "操作成功",
	SYSTEM_ERROR:               "服务器繁忙，稍后再试！",
	TIME_OUT:                   "网络请求超时！",
	FUNCTION_NOT_AVAILABLE:     "功能暂未开发",
	PARAMETER_ILLEGAL_DELETION: "参数非法或缺失！",

	DATA_NOT_EXIST: "数据不存在！",

	SERVER_ADDR_ERROR:            "邮件服务器地址是空的或错误的，请规范填写",
	EMAIL_ACCOUNT_ERROR:          "请填写你的邮箱账号",
	EMAIL_PASSWORD_ERROR:         "请填写你的邮箱密码",
	TIME_COUNT_ERROR:             "下载邮件的时间范围和下载指定数量的邮件只能二选一",
	EMAIL_DOWNLOAD_URL_ERROR:     "请填写你的附件下载目录，需要下载到哪里",
	EMAIL_LOGIN_ERROR:            "邮箱登录出错！请检查网络是否连接或邮箱信息是否填写正确无误",
	DOWNLOAD_URL_ERROR:           "检测到当前附件的下载路径是文件，请改为目录",
	DOWNLOAD_TIME_ERROR:          "下载邮件的时间格式错误",
	EMAIL_EMPTY_ERROR:            "邮件箱中未获取到邮件信息",
	FILE_TYPE_EMPTY_ERROR:        "请选择若附件存在同名文件时的操作",
	FILE_DOWNLOAD_CLASSIFY_ERROR: "请选择下载目录划分级别",
	GET_SYSTEM_PATH_ERROR:        "获取系统路径出现错误",
}

func GetErrMsg(code int64) string {
	if v, ok := ErrText[code]; ok {
		return v
	}
	return "服务器繁忙，稍后再试！"
}
