package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1") // Criação de grupos de APIS
	{
		v1.GET("/opening", handlers.GetOpeningHandler)
		v1.POST("/opening", handlers.PostOpeningHandler)
		v1.PUT("/opening", handlers.DeleteOpeningHandler)
		v1.DELETE("/opening", handlers.PutOpeningHandler)
	}

	v2 := router.Group("/api/v2") // Exemplo de teste de utilização de múltiplas versões ou grupos
	{
		v2.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "Desenvolvimento",
			})
		})
	}
}
