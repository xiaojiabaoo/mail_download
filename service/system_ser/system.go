package system_ser

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/pkg/errors"
	"io/ioutil"
	"mail_download/model/request_model"
	"mail_download/model/response_model"
	"mail_download/tools"
	customErr "mail_download/tools/error"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

func CheckUpdate(param request_model.CheckUpdateParam) (response_model.Version, error) {
	if param.Type == 1 {
		return response_model.Version{CurrentVersion: tools.Version}, nil
	}
	return checkUpdate("")
}

var (
	mu sync.Mutex
)

func checkUpdate(path string) (response_model.Version, error) {
	var (
		response          = response_model.Version{}
		url               = "https://gitee.com/xiaojiabaoo/mail_download.git"
		branch            = "main"
		err               error
		commit            = &object.Commit{}
		index, i          int
		newVersion        float64
		exist             bool
		appPath, tempPath string
	)
	mu.Lock()
	defer mu.Unlock()
	if path == "" {
		appPath, err = os.Getwd()
		if err != nil {
			return response, customErr.New(customErr.GET_APPPATH_ERROR, "")
		}
		tempPath = appPath + "\\Temp"
	} else {
		tempPath = path
	}
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
	// 克隆仓库到本地路径，使用浅克隆
	repo, err := git.PlainClone(tempPath, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		SingleBranch:  true,
		Depth:         1, // 只克隆最新的 commit
	})
	if err != nil {
		return response, errors.Wrap(err, "检测更新-拉取项目失败")
		os.Exit(1)
	}
	// 获取仓库的HEAD引用
	ref, err := repo.Head()
	if err != nil {
		return response, errors.Wrap(err, "检测更新-拉取项目失败")
		os.Exit(1)
	}
	// 获取最新的 commit 对象
	commit, err = repo.CommitObject(ref.Hash())
	if err != nil {
		fmt.Println("Error getting commit object:", err)
		os.Exit(1)
	}
	// 输出最新的 commit 信息
	//fmt.Println("Latest commit information:")
	//fmt.Println("Hash:", commit.Hash)
	//fmt.Println("Author:", commit.Author.Name)
	//fmt.Println("Email:", commit.Author.Email)
	//fmt.Println("Date:", commit.Author.When)
	//fmt.Println("Message:", commit.Message)
	index = strings.Index(commit.Message, "Version:")
	if index >= 0 {
		newVersion = tools.StringToFloat64(strings.ReplaceAll(commit.Message[index:], "Version:", ""))
	}
	i = strings.Index(commit.Message, ")：")
	if i == -1 {
		i = strings.Index(commit.Message, "):")
	}
	if i >= 0 {
		response.Describe = commit.Message[i+2 : index]
	}
	response.CurrentVersion = tools.Version
	response.NewVersion = newVersion
	if path == "" {
		os.RemoveAll(tempPath)
	}
	return response, nil
}

func Update() error {
	var (
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
	update, err = checkUpdate("")
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
	_, err = checkUpdate(path)
	if err != nil {
		return errors.Wrap(err, "拉取项目失败")
	}
	DeleteFile(path + "\\.git")
	return nil
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
