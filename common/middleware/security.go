package middleware

import (
	"fmt"

	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/response"
	"github.com/cuit9622/dms/common/response/errors"
	"github.com/cuit9622/dms/common/util/jwtUtil"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func SecurityMiddleWare(c *gin.Context) {
	permission := global.GLO_PERMISSON[c.Request.URL.String()]
	if permission == nil {
		return
	}

	tokenStr := c.GetHeader("token")
	id, err := jwtUtil.GetUserId(tokenStr)
	if err != nil {
		c.Abort()
		response.ErrorCode(c, errors.FORBIDEN)
		return
	}

	key := fmt.Sprintf("permissions:%d", id)
	str := global.GLO_REDIS.Get(key).Val()
	authorities := []string{}
	json.Unmarshal([]byte(str), &authorities)
	for _, auth := range authorities {
		if !permission.Contains(auth) {
			c.Abort()
			response.ErrorCode(c, errors.FORBIDEN)
			return
		}
	}

}
