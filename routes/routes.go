package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

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

	r.GET("/usuarios", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": usuarios,
		})
	})

	r.POST("/usuarios", func(c *gin.Context) {
		var nuevoUsuario Usuario

		if err := c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error al decodificar el JSON",
			})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Nombre y correo electronico son campos requeridos",
			})
			return
		}

		usuarios = append(usuarios, nuevoUsuario)

		c.JSON(http.StatusOK, gin.H{
			"message": "Usuario registrado",
			"data":    nuevoUsuario,
		})
	})

}

func fromFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"function": "from function",
	})
}
