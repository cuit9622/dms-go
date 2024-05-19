package initialize

import (
	"cuit9622/dms-common/response"
	"cuit9622/dms-common/response/errors"

	"github.com/gin-gonic/gin"
)

func initGin() *gin.Engine {
	g := gin.Default()
	g.NoRoute(func(c *gin.Context) {
		response.ErrorCode(c, errors.NOT_FOUND)
	})
	return g
}
