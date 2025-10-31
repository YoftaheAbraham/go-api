package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

type RouteResponse struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

func RegisterHandler(r *gin.Engine) {
	r.GET("/route", func(ctx *gin.Context) {
		response := controllers.HandleReq()
		ctx.JSON(200, RouteResponse{Message: response.Message, Path: ctx.Request.Host})
	})
}
