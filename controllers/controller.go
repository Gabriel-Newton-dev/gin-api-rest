package controllers

import (
	"net/http"

	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	// database.DB.Find(mais endereco de memoria), por se tratar de um slice com várias informacoes
	//temos que criar uma variável var alunos que recebe um slice da struct []models.Aluno
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "Seja bem-vindo" + nome + ", a nossa API que se comunica com banco de dados postgres?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}

func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado."})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"Deletado": "Usuário deletado do banco de dados com sucesso."})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro na função Edita Aluno"})
		return // caso tenha um erro ele irá sair da função.
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	// database.DB.Model(&aluno).UpdateColumns(aluno)
	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// a funcao edita alunos ela irá pegar o copor da nossa requisicao e mudar no BD
// Se tem o corpo a forma que usamos é o shouldBindJson para empacotar todo corpo da
//requisicao com base na nossa struct, com base no nosso endereco de memória da var aluno que criamos
// if err := .ShouldBindJSON()

func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": "CPF não localizado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func BuscaPorRg(c *gin.Context) {
	var aluno models.Aluno
	rg := c.Param("rg")
	database.DB.Where(&models.Aluno{RG: rg}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": "RG não localizado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
