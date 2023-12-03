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
		v1.GET("/alunos", handlers.GetAlunos)
		v1.GET("/aluno/:id", handlers.GetAlunoID)
		// v1.GET("/aluno/nome/:nome", handlers.GetAlunoNome)
		v1.POST("/aluno", handlers.JwtAuth, handlers.PostDB)
		v1.PUT("/aluno", handlers.PutDB)
		v1.DELETE("/aluno/:id", handlers.DeleteAluno)
		// v1.GET("/:param", handlers.TesteParam)
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
