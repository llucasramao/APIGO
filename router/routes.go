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
		v1.GET("/aluno", handlers.GetDB)
		v1.POST("/aluno", handlers.PostDB)
		v1.PUT("/aluno", handlers.DeleteDB)
		v1.DELETE("/aluno", handlers.PutDB)
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
