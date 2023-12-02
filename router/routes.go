package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize Handlers
	handlers.InitializeHandlers()

	v1 := router.Group("/api/v1") // Create APIs groups
	{
		v1.GET("/opening", handlers.GetOpeningHandler)
		v1.POST("/opening", handlers.PostOpeningHandler)
		v1.PUT("/opening", handlers.DeleteOpeningHandler)
		v1.DELETE("/opening", handlers.PutOpeningHandler)
		v1.GET("/:param", handlers.TesteParam)
	}

	v2 := router.Group("/api/v2") // Test example using multi version or groups
	{
		v2.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "Desenvolvimento",
			})
		})
	}
}
