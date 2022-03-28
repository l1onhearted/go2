package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "")
	})
	r.POST("/login",func(c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("")
	}
	r.Run(":9090")
}
