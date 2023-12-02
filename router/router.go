package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Create router using gin framework
	router := gin.Default()

	// Initialize routes
	InitializeRoutes(router)

	// Run HTTP server
	router.Run()
}
