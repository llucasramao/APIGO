package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Create Router
	router := gin.Default()

	// Initilize Routes
	InitializeRoutes(router)

	// Run Server
	router.Run()
}
