package main

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/Gabriel-Newton-dev/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{
			Nome: "Gabriel",
			CPF:  "107.879.987-09",
			RG:   "53535536-0",
		},
		{
			Nome: "Julia",
			CPF:  "145.908.879-09",
			RG:   "5143920-8",
		},
	}
	routes.HandleRequests()
}
