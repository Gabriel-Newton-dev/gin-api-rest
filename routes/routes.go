package routes

import "github.com/gin-gonic/gin"

func HandleRequest() {

	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos) //Temos uma rota r, e especificamos o verbo GET(pegar)irá retornar / colocando endpoint /alunos
	r.Run()
}
