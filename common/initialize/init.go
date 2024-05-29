package initialize

import (
	"net"

	"github.com/cuit9622/dms/common/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommon() net.Listener {
	initZap()
	initViper()
	ln := initListener()
	InitNacos()
	return ln
}

func InitSecurity() (*gin.Engine, net.Listener) {
	ln := InitCommon()
	g := initGin()
	InitRedis()
	g.Use(middleware.SecurityMiddleWare)
	return g, ln
}
