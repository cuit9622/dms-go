package response

import (
	"net/http"

	"github.com/cuit9622/dms/common/response/errors"

	"github.com/gin-gonic/gin"
)

type CommonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
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

func ErrorCode(c *gin.Context, errorCode *errors.ErrorCode) {
	c.JSON(http.StatusOK,
		CommonResult{
			Code: errorCode.Code,
			Msg:  errorCode.Msg,
			Data: nil,
		})
}
