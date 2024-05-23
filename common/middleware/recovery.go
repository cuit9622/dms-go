package middleware

import (
	"cuit9622/dms-common/response/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context, err any) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, errors.INTERNAL_SERVER_ERROR)
}
