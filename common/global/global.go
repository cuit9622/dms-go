package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GLO_VP  *viper.Viper
	GLO_LOG *zap.Logger
)
