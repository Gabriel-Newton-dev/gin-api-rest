package main

import "github.com/gin-gonic/gin"

// Essa funcao adaptei pelo que peguei na documentacao do Gin no GitHub
func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Gabriel Newton",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos) //Temos uma rota r, e especificamos o verbo GET(pegar)irá retornar / colocando endpoint /alunos
	r.Run()
}
