package routes

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.DisplayAllStudent)
	r.GET("/:nome", controllers.Salutation)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.GET("/alunos/:id", controllers.SearchStudentbyID)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.EditStudent)
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCpf)
	r.GET("/alunos/rg/:rg", controllers.SearchByRg)
	r.GET("/index", controllers.DisplayIndexPage)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}

// Na nossa rota r.GET("/alunos/:id", mais nome da funcao)
// colocamos ' : ' pq é a informacao que vai variar, assim como a funcao SAudacao acima
