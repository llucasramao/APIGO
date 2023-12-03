package handlers

import (
	"github.com/gin-gonic/gin"
)

func PutAluno(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"teste": "teste",
	})
}
