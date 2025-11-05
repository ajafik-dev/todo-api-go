package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := log.New(os.Stdout, "LOG:: ", log.LstdFlags)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Printf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	}
}
