package tools

import "sync"

var (
	LOG_LEVEL_INFO         = "INFO"         // 日志信息等级类型，日志存储级别：INFO
	LOG_LEVEL_SYSTEM_ERROR = "SYSTEM_ERROR" // 日志信息等级类型，日志存储级别：系统ERROR
	LOG_LEVEL_ERROR        = "ERROR"        // 日志信息等级类型，日志存储级别：逻辑ERROR
	Drive                  = []string{"C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:"}
)

const (
	CCL_RESULT_ALL_SUCCESS  = "全部成功"
	CCL_RESULT_ALL_FAIL     = "全部失败"
	CCL_RESULT_PART_SUCCESS = "部分成功"
	CCL_RESULT_PART_FAIL    = "部分失败"
	CCL_RESULT_FAIL         = "失败"
)

const (
	MAIL_RESULT_ALL_SUCCESS  = "全部成功"
	MAIL_RESULT_ALL_FAIL     = "全部失败"
	MAIL_RESULT_PART_SUCCESS = "部分成功"
	MAIL_RESULT_PART_FAIL    = "部分失败"
	MAIL_RESULT_FAIL         = "失败"
	MAIL_RESULT_NOT_COUNT    = "失败"
)

var (
	ProcessMap = sync.Map{}
)
