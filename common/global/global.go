package global

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GLO_VP        *viper.Viper
	GLO_LOG       *zap.Logger
	GLO_REDIS     *redis.Client
	GLO_PERMISSON map[string]*hashset.Set
	GLO_DB        *gorm.DB
)
