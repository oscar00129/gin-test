package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/params/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		message := fmt.Sprintf("Hola, %s", nombre)

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	r.GET("/func", fromFunction)
}

func fromFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"function": "from function",
	})
}
