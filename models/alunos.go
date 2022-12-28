package models

type Aluno struct {
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

// criei um slice de Aluno(como se fosse uma lista, pq terá vários)
var Alunos []Aluno
