package ginUtil

import (
	"github.com/cuit9622/dms/common/global"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gin-gonic/gin"
)

type GinUtil struct {
	router *gin.RouterGroup
}

func (i *GinUtil) GET(url string, handler gin.HandlerFunc, permission ...string) {
	i.router.GET(url, handler)
	addPermission(i.router.BasePath()+url, permission)
}
func (i *GinUtil) POST(url string, handler gin.HandlerFunc, permission ...string) {
	i.router.POST(url, handler)
	addPermission(i.router.BasePath()+url, permission)
}
func (i *GinUtil) PUT(url string, handler gin.HandlerFunc, permission ...string) {
	i.router.PUT(url, handler)
	addPermission(i.router.BasePath()+url, permission)
}
func (i *GinUtil) DELETE(url string, handler gin.HandlerFunc, permission ...string) {
	i.router.DELETE(url, handler)
	addPermission(i.router.BasePath()+url, permission)
}
func New(router *gin.RouterGroup) *GinUtil {
	return &GinUtil{
		router: router,
	}
}
func addPermission(url string, permission []string) {
	if len(permission) > 0 {
		set := hashset.New()
		for _, item := range permission {
			set.Add(item)
		}
		global.GLO_PERMISSON[url] = set
	}
}
