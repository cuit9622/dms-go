package initialize

import (
	"cuit9622/dms-common/response"

	"github.com/gin-gonic/gin"
)

func initGin() *gin.Engine {
	g := gin.Default()
	g.NoRoute(func(c *gin.Context) {
		response.Error(c, 404, "请求未找到")
	})
	return g
}
