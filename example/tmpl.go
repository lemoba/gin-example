package example

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website!",
	//	})
	//})
	//r.LoadHTMLGlob("templates/**/*")
	//r.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	//		"title": "Posts",
	//	})
	//})
	//r.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	//		"title": "Users",
	//	})
	//})

	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	r.LoadHTMLFiles("templates/raw.tmpl")
	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2023, 03, 16, 10, 10, 10, 10, time.UTC),
		})
	})
	r.Run(":8881") // 必须要加:
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

//func Post() {
//	var urls []string = []string{
//		"get-operation",
//		"get-realtime",
//		"get-output",
//		"get-pass",
//		"get-ng-station",
//		"get-yield",
//		"get-monthly",
//		"get-monthly-pass",
//	}
//	services.MulRequst(urls)
//}
