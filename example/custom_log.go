package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.New()

	r.Use(gin.LoggerWithFormatter(customLogger))
	r.GET("/log", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":8881")
}

func customLogger(params gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		params.ClientIP,
		params.TimeStamp.Format(time.RFC1123),
		params.Method,
		params.Path,
		params.Request.Proto,
		params.StatusCode,
		params.Latency,
		params.Request.UserAgent(),
		params.ErrorMessage,
	)
}
