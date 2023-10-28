package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"nepali-address-api/controllers"
	"net/http"
	"os" // Import the "os" package
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
	// Get the port from the environment variable, or use a default value (e.g., 8080)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)
}
