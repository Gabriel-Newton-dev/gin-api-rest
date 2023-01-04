package routes

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	// r.POST("/alunos", controllers.CriarNovoAluno)
	r.Run()
}
