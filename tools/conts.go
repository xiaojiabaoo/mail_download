package tools

import (
	"mail_download/tools/log/zap"
	"sync"
)

var (
	Maps   = sync.Map{}
	Logger *zap.ZapLog
	Drive  = []string{"C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:"}
)
