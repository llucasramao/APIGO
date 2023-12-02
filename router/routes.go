package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.GetOpeningHandler)
		v1.POST("/opening", handlers.PostOpeningHandler)
		v1.PUT("/opening", handlers.DeleteOpeningHandler)
		v1.DELETE("/opening", handlers.UpdateOpeningHandler)
	}
}
