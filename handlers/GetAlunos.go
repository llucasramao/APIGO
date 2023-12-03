package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func GetAlunos(ctx *gin.Context) {
	var alunos []models.Aluno

	config.DB.Find(&alunos)
	ctx.JSON(200, gin.H{
		"Response": alunos,
	})
}
