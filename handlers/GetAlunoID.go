package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func GetAlunoID(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	config.DB.First(&aluno, id)
	if aluno.ID == 0 {
		ctx.JSON(404, gin.H{
			"NotFound": "Aluno " + id + " Não encontrado",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Response": aluno,
	})
}
