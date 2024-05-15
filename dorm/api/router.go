package api

import "github.com/gin-gonic/gin"

func SetRouter(c *gin.Engine) {
	c.GET("test1", Test1)
	c.POST("test2", Test2)
	c.POST("test3", Test3)
}
