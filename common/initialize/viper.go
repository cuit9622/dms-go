package initialize

import (
	"fmt"

	"github.com/cuit9622/dms/common/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func initViper() {
	v := viper.New()
	v.BindEnv("MYSQL_SERVER", "MYSQL_SERVER")
	v.BindEnv("NACOS_SERVER", "NACOS_SERVER")
	v.BindEnv("REDIS_SERVER", "REDIS_SERVER")
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
	global.GLO_LOG.Info("Viper initialization complete")
}
