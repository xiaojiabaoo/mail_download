package tools

import "sync"

var (
	LOG_LEVEL_INFO         = "INFO"         // 日志信息等级类型，日志存储级别：INFO
	LOG_LEVEL_SYSTEM_ERROR = "SYSTEM_ERROR" // 日志信息等级类型，日志存储级别：系统ERROR
	LOG_LEVEL_ERROR        = "ERROR"        // 日志信息等级类型，日志存储级别：逻辑ERROR
	Drive                  = []string{"C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:"}
	Version                = 3.8
)

var (
	ProcessMap = sync.Map{}
	XlsxMap    = sync.Map{}
)
