package controllers

import (
	"net/http"

	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/Gabriel-Newton-dev/gin-api-rest/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Login
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind Json:" + err.Error(),
		})
		return
	}

	var user models.User
	dbError := db.Where("email = ?", p.Email).First(&user).Error
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find user"})
		return
	}
	if user.Password != services.SHA256Encoder(p.Password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"Error": "Invalid Credentials"})
		return
	}
	token, err := services.NewJwtService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token})
}
