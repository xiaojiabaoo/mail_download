package system_ser

import (
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/tools"
	customErr "mail_download/tools/error"
	"os"
)

func List(param request_model.SysListParam) ([]response_model.SystemList, error) {
	var (
		data []response_model.SystemList
		err  error
	)
	if param.Path == "" {
		//如果path没有，系统查不到需要查询那个目录，那就查询根目录（盘符）
		for _, partition := range tools.Drive {
			_, err = os.Stat(partition)
			if err == nil {
				data = append(data, response_model.SystemList{
					Path: partition + `\`,
					Name: partition,
				})
			}
		}
		return data, nil
	}
	files, err := os.ReadDir(param.Path)
	if err != nil {
		return nil, customErr.New(customErr.GET_SYSTEM_PATH_ERROR, "")
	}
	for _, v := range files {
		if v.IsDir() {
			data = append(data, response_model.SystemList{
				Path: param.Path + `\` + v.Name(),
				Name: v.Name(),
			})
		}
	}
	return data, nil
}
