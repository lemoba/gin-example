package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/logs", func(c *gin.Context) {
		fmt.Println("aaa")
		c.String(200, "write logs")
	})
	r.Run(":8881")
}
