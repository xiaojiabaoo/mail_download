package system_ser

import (
	"fmt"
	"github.com/pkg/errors"
	"mail_download/common/configx"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/tools"
	customErr "mail_download/tools/error"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
		if !v.IsDir() {
			continue
		}
		data = append(data, response_model.SystemList{
			Path: param.Path + `\` + v.Name(),
			Name: v.Name(),
		})
	}
	return data, nil
}

func Update() error {
	var (
		version      = configx.AppConfigData.Version
		url          = "https://gitee.com/xiaojiabaoo/mail_download.git"
		branch       = "main"
		err          error
		commit, path string
		appPath      string
		index        int
		newVersion   float64
	)
	appPath, err = os.Executable()
	if err != nil {
		return customErr.New(customErr.GET_APPPATH_ERROR, "")
	}
	// 获取上一级目录
	path = filepath.Join(filepath.Dir(appPath), "..")
	// 获取最新提交的信息，并从中截取版本号
	commit, err = GetCommit(url, branch)
	if err != nil {
		return err
	}
	index = strings.Index(commit, "Version:")
	if index == -1 {
		return customErr.New(customErr.COMMIT_VERSION_ERROR, "")
	}
	newVersion = tools.StringToFloat64(commit[index:])
	if newVersion == version {
		return customErr.New(customErr.IS_LATEST_VERSION, "")
	}
	if newVersion < version {
		return customErr.New(customErr.VERSION_NUMBER_ERROR, "")
	}
	path += fmt.Sprintf(`\\机器人 v%v`, newVersion)
	// 创建文件夹
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return errors.Wrap(err, "创建文件夹失败")
		}
	}
	cmd := exec.Command("git", "clone", "-b", branch, url, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return errors.Wrap(err, "拉取项目失败")
	}
	return nil
}

func GetCommit(url, branch string) (string, error) {
	cmd := exec.Command("git", "ls-remote", "--quiet", "--tags", "--heads", url, branch)
	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "获取远程存储库信息时出错")
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) <= 0 {
		return "", nil
	}
	fields := strings.Fields(lines[0])
	if len(fields) <= 1 {
		return "", nil
	}
	commitDescription, err := GetCommitDescribe(fields[0])
	if err != nil {
		return "", err
	}
	return commitDescription, nil
}

func GetCommitDescribe(hash string) (string, error) {
	cmd := exec.Command("git", "show", "-s", "--format=%B", hash)
	cmd.Stdout = nil // 设置为 nil 避免 Stdout 重复设置
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "获取提交描述时出错")
	}
	return strings.TrimSpace(string(output)), nil
}
