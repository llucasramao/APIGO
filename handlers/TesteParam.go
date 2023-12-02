package handlers

import "github.com/gin-gonic/gin"

func TesteParam(ctx *gin.Context) {
	nome := ctx.Params.ByName("param")
	ctx.JSON(200, gin.H{
		"status": "OK",
		"nome":   nome,
	})
}
