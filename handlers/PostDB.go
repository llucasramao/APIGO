package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func PostDB(ctx *gin.Context) {
	var aluno models.Aluno
	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		logger.Errf("Error ShouldBindJSON struct convert: %v", err)
		ctx.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	config.DB.Create(&aluno)
	ctx.JSON(200, gin.H{
		"Status":     "OK",
		"dataInsert": aluno,
	})
}
