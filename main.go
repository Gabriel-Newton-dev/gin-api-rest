package main

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/Gabriel-Newton-dev/gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gabriel", CPF: "123.456.123-31", RG: "12934234-1"},
		{Nome: "Guilheme", CPF: "345.456.432-43", RG: "32456432-1"},
	}
	routes.HandleRequest()
}
