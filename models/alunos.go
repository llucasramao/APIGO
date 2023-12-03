package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" binding:"required,min=3"`
	CPF  string `json:"cpf" binding:"required,min=3"`
	RG   string `json:"rg" binding:"required,min=3"`
}
