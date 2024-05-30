package initialize

import (
	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/util/httpClientUtil"
)

func InitHttpClient() {
	global.GLO_HTTP_CLIENT = httpClientUtil.New()
}
