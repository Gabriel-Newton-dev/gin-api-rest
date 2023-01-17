package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11"`
	RG   string `json:"rg" validate:"len=9"`
}

// criei um slice de Aluno(como se fosse uma lista, pq terá vários)
var Alunos []Aluno
