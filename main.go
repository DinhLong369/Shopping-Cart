package main

import (
	"Shopping-cart/config"
	"Shopping-cart/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connection()

	r := gin.Default()
	r1 := r.Group("/r1")
	routes.SetupRoute(r1)

	r.Run(":5000")
}
