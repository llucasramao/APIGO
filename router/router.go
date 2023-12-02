package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Criação da rota via framework Gin
	router := gin.Default()

	// Inicialização das rotas
	InitializeRoutes(router)

	// Execução do servidor HTTP
	router.Run()
}
