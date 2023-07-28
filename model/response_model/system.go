package response_model

type SystemList struct {
	Path     string `json:"path"`      // 目录名称
	FullPath string `json:"full_path"` // 目录全路径
	Type     int64  `json:"type"`      // 类型：1.目录；2文件
}
