package controllers

import (
	"net/http"

	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/Gabriel-Newton-dev/gin-api-rest/services"
	"github.com/gin-gonic/gin"
)

func CreateNewUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	user.Password = services.SHA256Encoder(user.Password)

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func DisplaysAllUser(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	// database.DB.Find(mais endereco de memoria), por se tratar de um slice com várias informacoes
	//temos que criar uma variável var alunos que recebe um slice da struct []models.Aluno
	c.JSON(200, alunos)
}

func DisplaysAllUserByID(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "User Not Found"})
	}

}
