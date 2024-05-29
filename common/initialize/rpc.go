package initialize

import (
	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/util/rpcUtil"
)

func InitRpc() {
	global.GLO_RPC = rpcUtil.New()
}
