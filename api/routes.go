package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(route *gin.Engine) {
	route.GET("/user/:id", GetUser)
}
