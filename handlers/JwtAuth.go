package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth(ctx *gin.Context) {
	if headerAuth := ctx.Request.Header.Get("Authorization"); !strings.HasPrefix(headerAuth, "Bearer ") {
		logger.Debug("teste")
		ctx.JSON(401, gin.H{
			"Status": "Unauthorized",
		})
		ctx.Abort()
	}
}
