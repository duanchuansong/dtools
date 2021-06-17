package config

import (
	"github.com/spf13/viper"
)

var C *viper.Viper

func InitConfig(path string) {
	C = viper.New()
	if path == "" {
		return
	}
	//设置读取的配置文件
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	C.SetConfigFile(path)

	if err := C.ReadInConfig(); err != nil {
		panic(err)
	}
}
