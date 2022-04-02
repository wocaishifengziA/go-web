package loggers

import (
	"fmt"
	"io"
	"os"
	"path"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var instance *logrus.Logger

var log = LogInstance()

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

func getFileInfo() *FileInfo {
	name := LOG_NAME
	dir := LOG_DIR
	return &FileInfo{
		Name:     name,
		Dir:      dir,
		FullPath: path.Join(dir, name),
	}
}

func InitLogWriteFile(isWriteFile bool) {
	if isWriteFile {
		fileInfo := getFileInfo()
		if err := os.MkdirAll(fileInfo.Dir, 0755); err != nil {
			fmt.Println(err.Error())
		}
		if _, err := os.Stat(fileInfo.FullPath); err != nil {
			if _, err := os.Create(fileInfo.FullPath); err != nil {
				fmt.Println(err.Error())
			}
		}
		writeToFile, err := os.OpenFile(fileInfo.FullPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("err", err)
		}
		writers := io.MultiWriter(os.Stdout, writeToFile)
		log.SetOutput(writers)
	}
}

func InitLogger() {
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&Formatter{})
	InitLogWriteFile(false)
}
