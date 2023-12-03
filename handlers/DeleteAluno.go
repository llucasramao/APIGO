package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func DeleteAluno(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")

	config.DB.First(&aluno, id)
	if aluno.ID == 0 {
		ctx.JSON(404, gin.H{
			"NotFound": "Aluno " + id + " Não encontrado",
		})
		return
	}

	config.DB.Delete(&aluno, id)
	ctx.JSON(200, gin.H{
		"status":  "deleted",
		"content": "Aluno " + id + " deletado",
	})
}
