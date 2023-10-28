package main

import (
	"nepali-address-api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  http.StatusOK,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "you are on the right place folks ;)",
			"status":  http.StatusOK,
		})
	})

	r.GET("/provinces", controllers.GetProvinces)
	r.GET("/districts/:province", controllers.GetDistrictsByProvince)

	r.Run() // listen and serve on
}
