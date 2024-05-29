package initialize

import (
	"fmt"

	"github.com/cuit9622/dms/common/global"

	"go.uber.org/zap"
)

func initZap() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Errorf("fatal error create logger: %s", err))
	}
	global.GLO_LOG = logger
	logger.Info("Zap initialization complete")
}
