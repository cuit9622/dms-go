package initialize

import (
	"cuit9622/dms-common/global"

	"github.com/go-redis/redis"
)

func InitRedis() {
	v := global.GLO_VP
	client := redis.NewClient(&redis.Options{
		Addr:     v.GetString("REDIS_SERVER") + ":6379", // redis地址
		Password: "",                                    // 密码
		DB:       0,                                     // 使用默认数据库
	})
	global.GLO_REDIS = client
	global.GLO_LOG.Info("Redis initialization complete")
}
