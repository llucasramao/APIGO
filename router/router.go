package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Create router using gin framework
	router := gin.Default()
	// router := gin.New()
	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))
	// router.Use(gin.Recovery())

	// Initialize routes
	InitializeRoutes(router)

	// Run HTTP server
	router.Run(":8080")
}
