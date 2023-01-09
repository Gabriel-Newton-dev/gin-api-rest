package routes

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("alunos/:id", controllers.EditaAluno)
	r.Run()
}

// Na nossa rota r.GET("/alunos/:id", mais nome da funcao)
// colocamos ' : ' pq é a informacao que vai variar, assim como a funcao SAudacao acima
