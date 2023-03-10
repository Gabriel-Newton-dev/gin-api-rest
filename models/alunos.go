package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
}

// criei um slice de Aluno(como se fosse uma lista, pq terá vários)
var Students []Aluno

func ValidateStudentData(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
