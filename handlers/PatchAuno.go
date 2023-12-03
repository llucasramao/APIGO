package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func PatchAluno(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	config.DB.Find(&aluno, id)
	if aluno.ID == 0 {
		ctx.JSON(404, gin.H{
			"NotFound": "Aluno " + id + " Não encontrado",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		logger.Errf("Error ShouldBindJSON struct convert: %v", err)
		ctx.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	config.DB.Model(&aluno).UpdateColumns(aluno)
	ctx.JSON(200, gin.H{
		"status":  "edited",
		"content": aluno,
	})
}
