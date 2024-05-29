package initialize

import (
	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/middleware"
	"github.com/cuit9622/dms/common/response"
	"github.com/cuit9622/dms/common/response/errors"

	"github.com/gin-gonic/gin"
)

func initGin() *gin.Engine {
	g := gin.New()
	g.Use(middleware.ZapLogger, gin.CustomRecovery(middleware.Recovery))
	g.NoRoute(func(c *gin.Context) {
		response.ErrorCode(c, errors.NOT_FOUND)
	})
	global.GLO_LOG.Info("Gin initialization complete")
	return g
}
