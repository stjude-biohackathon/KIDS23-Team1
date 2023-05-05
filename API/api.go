package main

import (
	// import control package from controller/control.go
	// import model package from model/model.go

	"fileFlipper/control"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrderPost struct {
	UserID     int    `json:"user_id"`
	SampleID   int    `json:"sample_id"`
	AnalysisID int    `json:"analysis_id"`
	OrderName  string `json:"order_name"`
}

func Get_orders_by_user(user_id string, db *gorm.DB) {
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
			user_name := ""
			user_department := ""

			user, err := control.Get_user_by_id(user_id, db)
			if err != nil {
				c.JSON(
					400,
					gin.H{"message": "bad request"},
				)
			} else {
				user_name = user.Name
				user_department = user.Department
			}

			fmt.Printf("User ID: %s, Name: %s\n", user_name, user_department)

			c.JSON(
				200,
				gin.H{
					"response": gin.H{
						"name":       user_name,
						"department": user_department,
					},
				},
			)
		}
	})

	router.GET("/samples/:sampleId", func(c *gin.Context) {
		sample_id := c.Param("sampleId")
		sample, err := control.Get_sample_and_related_files(sample_id, db)
		if err != nil {
			c.JSON(
				400,
				gin.H{"message": "bad request"},
			)
		} else {
			//extract sample
			selected_sample := sample
			//extract sampel file
			related_sample_files := sample.SampleFiles
			// return sample with related files using json
			c.JSON(
				200,
				gin.H{
					"response": gin.H{
						"sample": gin.H{
							"sample_id":    selected_sample.ID,
							"sample_name":  selected_sample.Sample_name,
							"sample_type":  selected_sample.Sample_type,
							"sample_state": selected_sample.Sample_state,
							"sample_files": related_sample_files,
						},
					},
				},
			)
		}
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
	})

	router.POST("/order", func(c *gin.Context) {
		var post OrderPost
		if err := c.BindJSON(&post); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		res := control.Submit_order(post.UserID, post.SampleID, post.AnalysisID, post.OrderName, db)

		print(res)

		c.JSON(200,
			gin.H{
				"message": res + " has started rehydrating process",
			})
	})

	router.GET("/analysis/all", func(c *gin.Context) {

		analyses, err := control.Get_analyses(db)

		if err != nil {
			c.JSON(
				400,
				gin.H{"message": "bad request"},
			)
		} else {
			c.JSON(
				200,
				gin.H{
					"response": gin.H{
						"analyses": analyses,
					},
				})
		}
	})

	router.Run(":8383")
}
