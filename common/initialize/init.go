package initialize

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	initViper()
	initZap()
	return initGin()
}
