package example

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/color.tmpl")

	r.GET("/myform", func(c *gin.Context) {
		c.HTML(200, "color.tmpl", gin.H{})
	})
	r.POST("/color", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		c.JSON(200, gin.H{
			"color": fakeForm.Colors,
		})
	})
	r.Run(":8881")
}
