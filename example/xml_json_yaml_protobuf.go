package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/some-json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "hi",
			"status": http.StatusOK,
		})
	})
	r.GET("/more-json", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "ranen"
		msg.Message = "你好, golang"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})
	r.GET("some-xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	r.GET("some-yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	//r.GET("some-protobuf", func(c *gin.Context) {
	//	reps := []int64{int64(1), int64(2)}
	//	label := "test"
	//
	//	c.ProtoBuf(http.StatusOK, data)
	//})
	r.Run(":8881")
}
