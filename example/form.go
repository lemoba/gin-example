package example

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		msg := c.PostForm("message")
		name := c.DefaultPostForm("name", "lemoba")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": msg,
			"name":    name,
		})
	})
	r.Run(":8888")
}
