package initialize

import (
	"2024Winter/app/api/global"
	"github.com/spf13/viper"
)

func SetupViper() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.SetConfigFile("./manifest/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper read config failed:" + err.Error())
	}

	err = viper.Unmarshal(&global.C)
	if err != nil {
		panic("viper unmarshal failed:" + err.Error())
	}
}
