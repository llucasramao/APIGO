package handlers

import "github.com/gin-gonic/gin"

func TesteParam(c *gin.Context) {
	nome := c.Params.ByName("param")
	c.JSON(200, gin.H{
		"status": "OK",
		"nome":   nome,
	})
}
