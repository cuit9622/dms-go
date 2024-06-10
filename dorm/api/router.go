package api

import (
	"github.com/cuit9622/dms/common/util/ginUtil"

	"github.com/gin-gonic/gin"
)

func SetRouter(c *gin.Engine) {
	router := ginUtil.New(c.Group(""))
	router.GET("test1", Test1)
	router.POST("test2", Test2, "dark")
	router.POST("test3", Test3, "homo", "fuck you")

	router.GET("dormBuilding", getDormBuildings)
	router.POST("dormBuilding", createDormBuilding)
	router.PUT("dormBuilding", updateDormBuilding)
	router.DELETE("dormBuilding/:id", deleteDormBuilding)

	router.GET("dorm", getDorms)
	router.PUT("dorm", updateDorm)
	router.DELETE("dorm/:id", deleteDorm)

	router.PUT("dormBed", updateDormBed)
	router.DELETE("dormBed/:id", deleteDormBed)
}
