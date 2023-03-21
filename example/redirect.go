package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://cn.bing.com")
	})

	r.GET("/foo", func(c *gin.Context) {
		c.String(200, "success")
	})

	r.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})
	//路由重定向，使用 HandleContext：
	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/foo"
		r.HandleContext(c)
	})
	r.Run(":8881")
}
