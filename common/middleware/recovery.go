package middleware

import (
	"net/http"

	"github.com/cuit9622/dms/common/response/errors"

	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context, err any) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, errors.INTERNAL_SERVER_ERROR)
}
