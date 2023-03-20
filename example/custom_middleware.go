package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前
		c.Next()

		// 请求后
		latency := time.Since(t)
		fmt.Println(latency)

		status := c.Writer.Status()

		fmt.Println(status)
	}
}
func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/middleware", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		fmt.Println(example)
	})
	r.Run(":8881")
}
