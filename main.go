package main

import (
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterHandler(r)

	r.Run(":9000")
}
