package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func CriarAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DeletarAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)

}

func TestVerificaEndpointSaudacao(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gabriel", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz":"Seja bem-vindo gabriel a nossa API"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockDaResposta, string(responseBody))

}

func TestListandoTodosAlunosHandler(t *testing.T) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder() // para armanezar todas as informacoes do corpo da nossa resposta
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code) // assertação para verificar se de fato estamos recebendo o status code que a gente espera.
	fmt.Println(response.Body)
	// t- teste, depois valor esperado - http.StatusOK, e o que vou receber é o status code da resposta
}
