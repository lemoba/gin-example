package example

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	router.Run(":8080")
}

func loginEndpoint(c *gin.Context)  {}
func submitEndpoint(c *gin.Context) {}
func readEndpoint(c *gin.Context)   {}
