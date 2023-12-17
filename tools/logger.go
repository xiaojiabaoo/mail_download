package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	mu sync.Mutex
)

func Logger(times int64, text, level string) {
	mu.Lock()
	defer mu.Unlock()
	if level == "" {
		level = LOG_LEVEL_INFO
	}
	dir := fmt.Sprintf("./logs")
	// 日志目录不存在会自己创建
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(filepath.Join(dir, fmt.Sprintf("%d.log", times)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println("关闭文件错误")
		}
	}(f)
	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("%s %s\n", level, text)
}
