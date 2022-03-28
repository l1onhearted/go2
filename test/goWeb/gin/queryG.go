package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Loger struct {
	User     string `form:"user" json:"user"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		name := c.DefaultQuery("name", "zyk")

		c.JSON(http.StatusOK, name)
	})
	r.GET("/:name/:age", func(c *gin.Context) {
		// name := c.DefaultQuery("name", "zyk")
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, name+age)
	})
	r.POST("/loginJSON",func(c *gin.Context){
		var loger Loger
		if err:=c.ShouldBind(&loger);err==nil{
			
		}
	})
	r.Run(":9090")
}
