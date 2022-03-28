package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//方法一:使用map
		// data := map[string]interface{}{
		// 	"name":    "zyk",
		// 	"message": "gogogogo",
		// 	"age":     18,
		// }
		//方法二:使用结构体
		data := gin.H{
			"name":    "zuk",
			"message": "gogog",
			"age":     "18",
		}
		//方法三:使用结构体
		var msg struct {
			Name    string `json:"name"`
			Message string `json:"message"`
			Age     int
		}
		msg.Name = "zyk"
		msg.Message = "gogog"
		msg.Age = 22
		c.JSON(http.StatusOK, data)
		c.JSON(http.StatusOK, msg)
	})
	r.Run(":9090")
}
