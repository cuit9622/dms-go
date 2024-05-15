package initialize

import (
	"cuit9622/dms-common/global"
	"fmt"

	"go.uber.org/zap"
)

func initZap() {
	// 初始化ZAP
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("fatal error create logger: %s", err))
	}
	defer logger.Sync() // zap底层有缓冲。在任何情况下执行 defer logger.Sync() 是一个很好的习惯
	global.GLO_LOG = logger
}
