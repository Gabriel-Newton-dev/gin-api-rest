package controllers

import (
	"net/http"

	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func DisplayAllStudent(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	// database.DB.Find(mais endereco de memoria), por se tratar de um slice com várias informacoes
	//temos que criar uma variável var alunos que recebe um slice da struct []models.Aluno
	c.JSON(200, alunos)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API says": "Welcome " + name + " to our API."})
}

func CreateNewStudent(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateStudentData(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}

func SearchStudentbyID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student Not Found"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeleteStudent(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"Deleted": "User successfully deleted from data base"})
}

func EditStudent(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error in function Edit Student"})
		return // caso tenha um erro ele irá sair da função.
	}
	if err := models.ValidateStudentData(&aluno); err != nil {
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

func SearchByCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": "Not Found CPF"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func SearchByRg(c *gin.Context) {
	var aluno models.Aluno
	rg := c.Param("rg")
	database.DB.Where(&models.Aluno{RG: rg}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not Found": "Not Found RG"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func CallViper() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

// funcao para rendereizar na tela index.html
func DisplayIndexPage(c *gin.Context) {
	var students []models.Aluno
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

// I created a func for when I have a route found

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
