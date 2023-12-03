package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func GetAlunoCpf(ctx *gin.Context) {
	var aluno models.Aluno
	cpf := ctx.Params.ByName("cpf")
	config.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		ctx.JSON(404, gin.H{
			"NotFound": "Aluno " + cpf + " Não encontrado",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "geted",
		"content": aluno,
	})
}
