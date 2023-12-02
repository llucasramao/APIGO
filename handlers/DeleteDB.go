package handlers

import (
	"github.com/gin-gonic/gin"
)

func DeleteDB(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"teste": "teste",
	})
}
