package tools

import (
	"fmt"
	"log"
	"mail_download/model"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func Logger(serial string, desc, resp, level string) {
	mu.Lock()
	defer mu.Unlock()
	if level == "" {
		level = LOG_LEVEL_INFO
	}
	if serial == "" {
		return
	}
	dir := "./logs/" + Serial(serial)
	// 日志目录不存在会自己创建
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(filepath.Join(dir, fmt.Sprintf("%s.log", serial)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println("关闭文件错误")
		}
	}(f)
	Excel(serial, model.XlsxData{Level: level, Describe: desc, Response: resp})
	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("%s %s %s\n", level, desc, resp)
}

func Excel(serial string, data model.XlsxData) {
	var (
		xlsx  = make([]model.XlsxData, 0, 0)
		now   = time.Now()
		ok    bool
		value any
	)
	value, ok = XlsxMap.Load(serial)
	if ok {
		xlsx = value.([]model.XlsxData)
	}
	xlsx = append(xlsx, model.XlsxData{
		Time:     now.Format("2006-01-02 15:04:05"),
		Level:    data.Level,
		Describe: data.Describe,
		Response: data.Response,
	})
	XlsxMap.Store(serial, xlsx)
}
