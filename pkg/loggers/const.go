package loggers

import "time"

const (
	defaultLogFormat       = "[%lvl%] %time% | %msg%"
	defaultLogPathFormat   = "[%lvl%] %time% %path% | %msg%"
	defaultTimestampFormat = time.RFC3339
)
