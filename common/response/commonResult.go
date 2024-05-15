package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResult struct {
	Code int
	Msg  string
	Data any
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, CommonResult{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK,
		CommonResult{
			Code: code,
			Msg:  msg,
			Data: nil,
		})
}
