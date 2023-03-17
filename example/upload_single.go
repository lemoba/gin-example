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

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./upload/" + file.Filename
		// 上传文件至指定的完整文件
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run(":8881")
}
