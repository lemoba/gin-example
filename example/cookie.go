package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "127.0.0.1", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
		c.String(200, "success")
	})
	r.Run(":8881")
}
