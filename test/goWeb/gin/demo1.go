package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello gin",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", sayHello)
	r.GET("/hello", sayHello)
	r.GET("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"meth":"GET",
		})
	})
	r.Run(":9090")
}
