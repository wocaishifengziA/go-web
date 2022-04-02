package loggers

import (
	"fmt"
	"io"
	"os"
	"path"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/wocaishifengziA/go-web/pkg/configs"
)

var once sync.Once
var instance *logrus.Logger

var logger = LogInstance()

func LogInstance() *logrus.Logger {
	once.Do(func() {
		instance = logrus.New()
	})
	return instance
}

type FileInfo struct {
	Name     string
	Dir      string
	FullPath string
}

func getFileInfo(c configs.Log) *FileInfo {
	name := c.FileName
	dir := c.FileDir
	return &FileInfo{
		Name:     name,
		Dir:      dir,
		FullPath: path.Join(dir, name),
	}
}

func InitLogWriteFile(c configs.Log) {
	if c.WriteFile {
		fileInfo := getFileInfo(c)
		if err := os.MkdirAll(fileInfo.Dir, 0755); err != nil {
			fmt.Println(err.Error())
			return
		}
		if _, err := os.Stat(fileInfo.FullPath); err != nil {
			if _, err := os.Create(fileInfo.FullPath); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		writeToFile, err := os.OpenFile(fileInfo.FullPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		writers := io.MultiWriter(os.Stdout, writeToFile)
		logger.SetOutput(writers)
	}
}

func InitLogger(c configs.Log) {
	logger.SetLevel(c.Level)
	logger.SetReportCaller(true)
	logger.SetFormatter(&Formatter{})
	InitLogWriteFile(c)
}
