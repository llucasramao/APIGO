package handlers

import (
	"github.com/gin-gonic/gin"
)

func PutDB(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"teste": "teste",
	})
}
