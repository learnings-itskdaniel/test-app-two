package main

import (
	_ "github.com/diegoluisi/test-app-two/config"
	"github.com/diegoluisi/test-app-two/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Hello world for your app test-app-two",
		})
	})
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)
	server.Run()
}
