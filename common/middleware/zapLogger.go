package middleware

import (
	"time"

	"github.com/cuit9622/dms/common/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	c.Next()

	cost := time.Since(start)
	global.GLO_LOG.Info(c.Request.Method,
		zap.String("path", path),
		zap.String("query", query),
		zap.Int("status", c.Writer.Status()),
		zap.String("ip", c.ClientIP()),
		zap.Duration("cost", cost),
	)
}
