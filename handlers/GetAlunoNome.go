package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func GetAlunoNome(ctx *gin.Context) {
	var aluno models.Aluno
	nome := ctx.Params.ByName("nome")
	config.DB.Delete(nome)
	if aluno.ID == 0 {
		ctx.JSON(404, gin.H{
			"NotFound": "Aluno " + nome + " Não encontrado",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Response": aluno,
	})
}
