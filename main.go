package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oscar00129/gin-test/routes"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
