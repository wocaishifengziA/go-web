package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	// 设置默认值
	v.SetDefault("setdefault", "abc")
	// 显示设置
	v.Set("age", 1)
	// 读取配置文件
	v.SetConfigFile("./config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config err: %s", err.Error()))
	}
	fmt.Println(v.AllSettings())

	// 写入配置文件
	err = v.WriteConfig()
	if err != nil {
		panic(fmt.Errorf("write config err: %s", err.Error()))
	}

	// 监听和重新加载
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(v.AllSettings())
	})
	v.WatchConfig()

	// 环境变量
	os.Setenv("SPF_ID", "13")
	os.Setenv("ENVA", "enva")
	// Env prefix
	v.SetEnvPrefix("spf")
	v.BindEnv("id")
	fmt.Println("env:", v.Get("id"))
	// Enc not prefix
	v.BindEnv("uu", "ENVA")
	fmt.Println("env uu:", v.Get("uu"))

	
}

func Wait() {
	var quit chan struct{}
	<-quit
}
