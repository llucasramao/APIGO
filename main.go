package main

import (
	"github.com/llucasramao/APIGO/config"
	"github.com/llucasramao/APIGO/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("Erro na conexão: %v", err)
		return
	}

	router.Initialize()
}
