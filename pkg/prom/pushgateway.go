package prom

import (
	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

const (
	PushGatewayAddr = "http://192.168.2.125:9091/"
)

var (
	InferCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_num",
			ConstLabels: prometheus.Labels{
				"serviceId":  "ab44",
				"versionNum": "v11",
			},
		},
		[]string{"code", "method"},
	)
)

func init() {
	log.Println("Init prometheus MustRegister InferCounter")
	prometheus.MustRegister(InferCounter)
}

func UploadPrometheus() {
	InferCounter.WithLabelValues("aaa", "111").Inc()
	InferCounter.WithLabelValues("aaa", "111").Inc()
	InferCounter.WithLabelValues("bbb", "222").Inc()
	if err := push.New(PushGatewayAddr, "infer_data").
		Collector(InferCounter).
		// Grouping("service", "xsddada-xsd898"+"@"+"v1").
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
