package api

import (
	"cuit9622/dms-common/util/ginUtil"

	"github.com/gin-gonic/gin"
)

func SetRouter(c *gin.Engine) {
	router := ginUtil.New(c.Group(""))
	router.GET("test1", Test1, "van")
	router.POST("test2", Test2, "dark")
	router.POST("test3", Test3, "homo", "fuck you")
}
