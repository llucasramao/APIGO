package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/models"
)

func DeleteAluno(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	config.DB.Delete(&aluno, id)

	ctx.JSON(200, gin.H{
		"Deleted": "Aluno " + id + " deletado",
	})
}
