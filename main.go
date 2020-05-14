package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-ssilva/api.go.gingonic/middlewares"
	"github.com/m-ssilva/api.go.gingonic/services"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.ErrorHandle())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": true,
		})
	})

	r.POST("/register", services.RegisterService)

	r.Run(":3000")
}
