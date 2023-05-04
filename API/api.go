package main

import (
	"API/control"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/:userId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		control.Get_orders_by_user()
	})

	router.GET("/samples", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		control.Get_samples()
	}

	router.GET("/orders/:userId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		control.Get_orders_by_user_id()
	})

	router.POST("/orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		control.Submit_order()
	}

	router.Run(":8383")
}
