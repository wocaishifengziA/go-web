package loggers

import "time"

const (
	defaultLogFormat       = "[%lvl%] %time% | %msg%"
	defaultLogPathFormat   = "[%lvl%] %time% %path% | %msg%"
	defaultTimestampFormat = time.RFC3339

	LOG_NAME = "xx.log"
	LOG_DIR  = "/data/apharbor/logs"
)
