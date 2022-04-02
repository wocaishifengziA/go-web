package configs

import "github.com/sirupsen/logrus"

type AppConfig struct {
	DebugModel bool
	Log        Log
}

type Log struct {
	Level      logrus.Level
	WriteFile  bool
	CallReport bool
	FileDir    string
	FileName   string
}
