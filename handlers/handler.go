package handlers

import (
	"github.com/llucasramao/APIGO/config"
)

var (
	logger *config.Logger
)

func InitializeHandlers() {
	logger = config.GetLogger("handlers")
}
