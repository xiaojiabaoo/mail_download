package system_ser

import (
	"fmt"
	"github.com/pkg/errors"
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

func CheckUpdate() (response_model.Version, error) {
	return checkUpdate()
}

func checkUpdate() (response_model.Version, error) {
	var (
		response          = response_model.Version{}
		url               = "https://gitee.com/xiaojiabaoo/mail_download.git"
		branch            = "main"
		err               error
		commit            string
		index, i          int
		newVersion        float64
		exist             bool
		appPath, tempPath string
	)
	appPath, err = os.Getwd()
	if err != nil {
		return response, customErr.New(customErr.GET_APPPATH_ERROR, "")
	}
	tempPath = appPath + "\\Temp"
	exist, err = DirExist(tempPath)
	if err != nil {
		return response, err
	}
	if exist {
		os.RemoveAll(tempPath)
	}
	// 创建文件夹
	err = os.Mkdir(tempPath, os.ModePerm)
	if err != nil && !strings.Contains(err.Error(), "file already exists") {
		return response, errors.Wrap(err, "检测更新-创建文件夹失败")
	}
	cmd := exec.Command("git", "clone", "-b", branch, url, tempPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return response, errors.Wrap(err, "检测更新-拉取项目失败")
	}
	// 获取最新提交的信息，并从中截取版本号
	commit, err = GetCommit(url, branch, appPath, tempPath)
	if err != nil {
		return response, err
	}
	index = strings.Index(commit, "Version:")
	if index >= 0 {
		newVersion = tools.StringToFloat64(strings.ReplaceAll(commit[index:], "Version:", ""))
	}
	i = strings.Index(commit, ")：")
	if i == -1 {
		i = strings.Index(commit, "):")
	}
	if i >= 0 {
		response.Describe = commit[i+2 : index]
	}
	response.CurrentVersion = tools.Version
	response.NewVersion = newVersion
	os.RemoveAll(tempPath)
	return response, nil
}

func Update() error {
	var (
		url     = "https://gitee.com/xiaojiabaoo/mail_download.git"
		branch  = "main"
		err     error
		path    string
		appPath string
		update  response_model.Version
		files   []os.DirEntry
	)
	appPath, err = os.Getwd()
	if err != nil {
		return customErr.New(customErr.GET_APPPATH_ERROR, "")
	}
	// 获取上一级目录
	path = filepath.Join(filepath.Dir(appPath), ".")
	update, err = checkUpdate()
	if err != nil {
		return err
	}
	if update.NewVersion == 0 {
		return customErr.New(customErr.COMMIT_VERSION_ERROR, "")
	}
	if update.NewVersion == update.CurrentVersion {
		return customErr.New(customErr.IS_LATEST_VERSION, "")
	}
	if update.NewVersion < update.CurrentVersion {
		return customErr.New(customErr.VERSION_NUMBER_ERROR, "")
	}
	files, err = os.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "获取程序父级目录下的目录信息出现问题")
	}
	for _, file := range files {
		if file.IsDir() && strings.Contains(file.Name(), fmt.Sprintf(`Version %v`, update.NewVersion)) {
			version := tools.StringToFloat64(strings.ReplaceAll(file.Name(), "Version ", ""))
			if version >= update.NewVersion {
				return customErr.New(customErr.VERSION_ERROR,
					fmt.Sprintf(`更新的版本为：Version %v 系统检测到你的电脑中已经存在了相同或更高的版本，为保证更新正常进行，请先删除相同和高版本的程序后重新尝试`, update.NewVersion))
			}
		}
	}
	path += fmt.Sprintf(`\Version %v`, update.NewVersion)
	// 创建文件夹
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "创建文件夹失败")
	}
	cmd := exec.Command("git", "clone", "-b", branch, url, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return errors.Wrap(err, "拉取项目失败")
	}
	DeleteFile(path + "\\.git")
	return nil
}

func GetCommit(url, branch, appPath, tempPath string) (string, error) {
	err := os.Chdir(tempPath)
	if err != nil {
		return "", err
	}
	defer func() error {
		err = os.Chdir(appPath)
		if err != nil {
			return err
		}
		return nil
	}()
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

func DeleteFile(filePath string) error {
	return os.RemoveAll(filePath)
}

func DirExist(path string) (bool, error) {
	// 使用 os.Stat 获取目录信息
	_, err := os.Stat(path)
	// 判断目录是否存在
	switch {
	case err == nil:
		return true, nil
	case os.IsNotExist(err):
		return false, nil
	default:
		return false, errors.Wrap(err, "检测目录信息错误")
	}
}
