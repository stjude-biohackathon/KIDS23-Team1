package main

import (
	// import control package from controller/control.go

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Get_orders_by_user(user_id string, db *gorm.DB) {
	print("Hello World fronm order controller!")
}

func Get_user_by_id(user_id string, db *gorm.DB) {
	print("Hello World fronm order controller!")
}

func main() {
	dsn := "host=localhost user=fileadmin password=test1234 dbname=fileflipperdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()

	router.GET("/user/:userId", func(c *gin.Context) {
		user_id := c.Param("userId")
		if user_id == "" {
			c.JSON(
				400,
				gin.H{"message": "bad request"},
			)
		} else {
			Get_user_by_id(user_id, db)
			c.JSON(
				200,
				gin.H{
					"body": "waiting for list",
				},
			)
		}
	})

	router.GET("/samples", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		// control.Get_samples()
	})

	router.GET("/orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		// control.Get_sample_and_related_files()
	})

	router.GET("/orders/:userId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		// control.Get_orders_by_user_id()
	})

	router.POST("/orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
		// control.Submit_order()
	})

	router.Run(":8383")
}
