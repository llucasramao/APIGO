package handlers

import (
	"github.com/llucasramao/APIGO/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandlers() {
	logger = config.GetLogger("handlers")
	db = config.GetSQLite()
}
