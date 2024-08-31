package system_ser

import (
	"io/ioutil"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/tools"
	"os"
	"path/filepath"
	"strings"
)

func List(param request_model.SysListParam) (response_model.SystemDirs, error) {
	var (
		data = response_model.SystemDirs{}
		err  error
	)
	data.Dirs = make([]response_model.SystemList, 0)
	dir := filepath.Dir(param.Path)
	if dir == "." {
		dir = ""
	}
	data.CurrentPath = param.Path
	data.ParentPath = dir
	// 如果去掉后路径等于输入路径，说明是根目录或者只有一层目录
	if dir == param.Path {
		data.ParentPath = ""
	}
	// 如果路径为空，返回所有盘符
	if param.Path == "" {
		for _, v := range GetAllDrives() {
			data.Dirs = append(data.Dirs, response_model.SystemList{
				Path: v,
				Name: v,
			})
		}
		return data, nil
	}
	// 检查路径是否存在
	info, err := os.Stat(param.Path)
	if err != nil {
		return data, err
	}
	// 如果路径是盘符，且没有斜杠（兼容传入“D:”的情况）
	if info.IsDir() && strings.HasSuffix(param.Path, ":") {
		param.Path += "\\"
	}
	// 读取路径下的所有文件和目录
	files, err := ioutil.ReadDir(param.Path)
	if err != nil {
		return data, err
	}
	// 过滤出目录
	var directories []string
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, filepath.Join(param.Path, file.Name()))
		}
	}
	for _, path := range directories {
		var (
			split = strings.Split(path, "\\")
			name  string
		)
		if len(split) > 0 {
			name = split[len(split)-1]
		}
		if _, ok := tools.DirMap[name]; ok {
			continue
		}
		data.Dirs = append(data.Dirs, response_model.SystemList{
			Path: path,
			Name: name,
		})
	}
	return data, nil
}

// 获取所有盘符
func GetAllDrives() []string {
	drives := []string{}
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		drivePath := string(drive) + ":\\"
		if _, err := os.Stat(drivePath); !os.IsNotExist(err) {
			drives = append(drives, drivePath)
		}
	}
	return drives
}

func Version(param request_model.VersionParam) response_model.Version {
	return response_model.Version{CurrentVersion: tools.Version}
}
