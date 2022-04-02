package loggers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wocaishifengziA/go-web/pkg/configs"
)

type Formatter struct {
	TimestampFormat string
	LogFormat       string
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat

		if configs.Config.Log.CallReport {
			output = defaultLogPathFormat
		}
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)
	time.Sleep(time.Duration(3) * time.Second)
	if configs.Config.Log.CallReport {
		output = strings.Replace(output, "%path%", fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line), 1)
	}

	if len(entry.Data) > 0 {
		fields, _ := json.Marshal(entry.Data)
		output = output + " | " + string(fields)
	}

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1) + "\n"

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil

}
