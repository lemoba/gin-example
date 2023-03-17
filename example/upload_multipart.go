package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("templates/upload.tmpl")

	r.GET("/upload-html", func(c *gin.Context) {
		c.HTML(200, "upload.tmpl", gin.H{})
	})

	r.MaxMultipartMemory = 8 << 2
	r.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		dst := "./upload/"
		for _, file := range files {
			log.Println(file, file.Filename)
			c.SaveUploadedFile(file, dst+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	r.Run(":8881")
}
