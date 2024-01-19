package request_model

type SysListParam struct {
	Path string `form:"path"`
}

type CheckUpdateParam struct {
	Type int64 `json:"type"` //1.表示仅检测当前版本；0.表示检测当前版本和新版本；
}
