package routes

import "github.com/gin-gonic/gin"

type RouteResponse struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

func RegisterHandler(r *gin.Engine) {
	r.GET("/route", func(ctx *gin.Context) {
		ctx.JSON(200, RouteResponse{Message: "From route handler", Path: ctx.Request.Host})
	})
}
