package initialize

import (
	"cuit9622/dms-common/middleware"
	"net"

	"github.com/gin-gonic/gin"
)

func InitCommon() (*gin.Engine, net.Listener) {
	initZap()
	initViper()
	g, ln := initGin()
	InitNacos()
	return g, ln
}
func InitSecurity() (*gin.Engine, net.Listener) {
	g, ln := InitCommon()
	InitRedis()
	g.Use(middleware.SecurityMiddleWare)
	return g, ln
}
