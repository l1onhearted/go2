package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "test user index",
			})
		})
		// userGroup.GET("/login", func(c *gin.Context) {...})
		// userGroup.POST("/login", func(c *gin.Context) {...})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "test shop index",
			})
		})
		// shopGroup.GET("/cart", func(c *gin.Context) {...})
		// shopGroup.POST("/checkout", func(c *gin.Context) {...})
	}
	r.Run(":9090")
}
