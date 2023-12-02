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

	router.Initialize()
}
