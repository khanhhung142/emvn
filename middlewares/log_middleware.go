package middlewares

import (
	"emvn/consts"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LogMiddleware is a middleware to log the request
// When error occurs, we set the error to the gin context, so that we can log the error in this middleware
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()

		errStr := ""
		if err, ok := c.Get(consts.GinErrorKey); ok {
			errStr = " | " + err.(consts.CustomError).Error()
		}

		latency := time.Since(t)

		statusCode := c.Writer.Status()
		stdOut := fmt.Sprintf("[GF] %v | %3d | %v | %v | %v | %v %#v %v\n",
			t.Format("2006/01/02 - 15:04:05 -07"),
			statusCode,
			latency,
			c.ClientIP(),
			c.GetHeader("User-Agent"),
			c.Request.Method,
			c.Request.RequestURI,
			errStr,
		)

		log.Println(stdOut)
	}
}
