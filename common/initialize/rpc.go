package initialize

import (
	"cuit9622/dms-common/global"
	"cuit9622/dms-common/util/rpcUtil"
)

func InitRpc() {
	global.GLO_RPC = rpcUtil.New(global.GLO_LOG, global.GLO_NACOS)
}
