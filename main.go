package main

import (
	"github.com/wocaishifengziA/go-web/pkg/configs"
	"github.com/wocaishifengziA/go-web/pkg/loggers"
)

func main() {
	configs.InitConfig("./conf/config.yaml")
	loggers.InitLogger(configs.Config.Log)
	loggers.LogInstance().Infoln("ok")
}
