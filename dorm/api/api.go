package api

import (
	"cuit9622/dms-common/response"
	"cuit9622/dms-common/util/jwtUtil"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func Test1(c *gin.Context) {
	response.Success(c, "Test1")
}

func Test2(c *gin.Context) {
	test := &Test{}
	c.ShouldBindBodyWithJSON(test)
	response.Success(c, test)
}

func Test3(c *gin.Context) {
	tokenStr := c.GetHeader("token")
	id, _ := jwtUtil.GetUserId(tokenStr)
	response.Success(c, fmt.Sprintf("Test2 %d", id))
}
