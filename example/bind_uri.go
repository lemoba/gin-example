package example

import "github.com/gin-gonic/gin"

type Persons struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Persons
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSONP(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	r.Run(":8881")
}
