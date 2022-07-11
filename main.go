package main

import (
	"log"

	"github.com/wocaishifengziA/go-web/pkg/prom"
)

func main() {
	// configs.InitConfig("./conf/config.yaml")
	// loggers.InitLogger(configs.Config.Log)
	// loggers.LogInstance().Infoln("ok")
	// rabbitmq.DoBlock()

	// promethus pushgateway
	log.Println("----------------")
	
	prom.UploadPrometheus()
	// prom.UploadPrometheus()
	// prom.UploadPrometheus()
	for{}
}
