package initialize

import (
	"cuit9622/dms-common/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommon() *gin.Engine {
	initViper()
	initZap()
	return initGin()
}
func InitSecurity() *gin.Engine {
	g := InitCommon()
	InitRedis()
	g.Use(middleware.SecurityMiddleWare)
	return g
}
