package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *AppConfig

func InitConfig(configFile string) {
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error read config file: %s ", err))
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config file: %s ", err))
	}
	fmt.Printf("config = %+v\n", *Config)
}

