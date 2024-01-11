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
	NO_RECORDS_FOUND           = 10005
	OPERATING_SYSTEM_ERROR     = 10006

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
	NUMBER_DOWNLOADS_100_ERROR   = 30013
	MAIL_BOX_EMPTY_ERROR         = 30014

	CLL_LOGIN_ERROR             = 40001
	INVOICENO_EMPTY_ERROR       = 40002
	INVOICENO_ERROR             = 40003
	DIRECTORY_FILE_ERROR        = 40004
	READING_FILE_ERROR          = 40005
	OPEN_FILE_ERROR             = 40006
	EMAIL_FORMAT_ERROR          = 40007
	TOO_MANY_FILES              = 40008
	ACCOUNT_TWO_PROCESSES_ERROR = 40009
)

var ErrText = map[int64]string{
	OPERATION_SUCCESS:          "操作成功",
	SYSTEM_ERROR:               "服务器繁忙，稍后再试！",
	TIME_OUT:                   "网络请求超时！",
	FUNCTION_NOT_AVAILABLE:     "功能暂未开发",
	PARAMETER_ILLEGAL_DELETION: "参数非法或缺失！",
	NO_RECORDS_FOUND:           "没有查询到数据",
	OPERATING_SYSTEM_ERROR:     "本程序暂不支持非Windows操作系统的设备运行，请尝试更换其他Windows设备使用",

	DATA_NOT_EXIST: "数据不存在！",

	SERVER_ADDR_ERROR:            "邮箱服务器填写错误或无法连接",
	EMAIL_ACCOUNT_ERROR:          "请填写你的邮箱账号",
	EMAIL_PASSWORD_ERROR:         "请填写你的邮箱密码",
	TIME_COUNT_ERROR:             "下载邮件的时间范围和下载指定数量的邮件必须要填写一个",
	EMAIL_DOWNLOAD_URL_ERROR:     "请填写你的附件下载目录，需要下载到哪里",
	EMAIL_LOGIN_ERROR:            "邮箱账号密码错误或无法登录",
	DOWNLOAD_URL_ERROR:           "检测到当前附件的下载路径是文件，请改为目录",
	DOWNLOAD_TIME_ERROR:          "下载邮件的时间不合法",
	EMAIL_EMPTY_ERROR:            "邮件箱中未获取到邮件信息",
	FILE_TYPE_EMPTY_ERROR:        "请选择若附件存在同名文件时的操作",
	FILE_DOWNLOAD_CLASSIFY_ERROR: "请选择下载目录划分级别",
	GET_SYSTEM_PATH_ERROR:        "获取系统路径出现错误",
	NUMBER_DOWNLOADS_100_ERROR:   "根据数量下载时不能超过100封邮件",
	MAIL_BOX_EMPTY_ERROR:         "请选需要下载的择收件箱",

	CLL_LOGIN_ERROR:             "账号或密码错误，请重新填写",
	INVOICENO_EMPTY_ERROR:       "未识别到单号或是单号错误",
	INVOICENO_ERROR:             "单号无法识别为空运或是海运",
	DIRECTORY_FILE_ERROR:        "请确保该目录下只含有pdf文件",
	READING_FILE_ERROR:          "读取目录文件失败，请检查目录地址是否填写正确",
	OPEN_FILE_ERROR:             "打开%s文件失败，请检查文件是否被顺坏",
	EMAIL_FORMAT_ERROR:          "请填写正确的邮箱格式",
	TOO_MANY_FILES:              "单次操作的PDF数量不能超过500个，请减少文件数量为<=500",
	ACCOUNT_TWO_PROCESSES_ERROR: "一个账号不能同时操作多个流程，请等待当前流程操作完成",
}

func GetErrMsg(code int64) string {
	if v, ok := ErrText[code]; ok {
		return v
	}
	return "服务器繁忙，稍后再试！"
}
