package response_model

type SystemDirs struct {
	ParentPath  string       `json:"parent_path"`  // 上一级目录路径
	CurrentPath string       `json:"current_path"` // 当前目录路径
	Dirs        []SystemList `json:"dirs"`
}

type SystemList struct {
	Path string `json:"path"` // 目录路径
	Name string `json:"name"` // 目录名称
}

type Version struct {
	CurrentVersion string `json:"current_version"`
}
