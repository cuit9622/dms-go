package global

import (
	"github.com/cuit9622/dms/common/interfaces"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-redis/redis"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Info struct {
	IP   string
	Port uint64
	Name string
}

var (
	GLO_INFO      Info
	GLO_VP        *viper.Viper
	GLO_LOG       *zap.Logger
	GLO_REDIS     *redis.Client
	GLO_PERMISSON map[string]*hashset.Set = map[string]*hashset.Set{}
	GLO_DB        *gorm.DB
	GLO_NACOS     naming_client.INamingClient
	GLO_RPC       interfaces.RpcUtil
)
