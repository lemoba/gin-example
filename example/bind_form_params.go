package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()

	r.GET("/testing", bindPage)
	r.Run(":8881")
}

func bindPage(c *gin.Context) {
	var user User

	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
	if c.ShouldBind(&user) == nil {
		fmt.Println(user.Name)
		fmt.Println(user.Address)
		fmt.Println(user.Birthday)
	}
	c.String(200, "success")
}
