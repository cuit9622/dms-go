package initialize

import (
	"cuit9622/dms-common/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() {
	// 初始化Viper
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
	global.GLO_VP = v

	// 初始化ZAP
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("fatal error create logger: %s", err))
	}
	defer logger.Sync() // zap底层有缓冲。在任何情况下执行 defer logger.Sync() 是一个很好的习惯
	global.GLO_LOG = logger
}
