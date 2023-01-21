package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
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
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder() // para armanezar todas as informacoes do corpo da nossa resposta
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code) // assertação para verificar se de fato estamos recebendo o status code que a gente espera.
	// t- teste, depois valor esperado - http.StatusOK, e o que vou receber é o status code da resposta
}

// func TestVerificaEdndpointCriaAluno(t *testing.T) {
// 	r := SetupDasRotasDeTeste()
// 	r.POST("/alunos", controllers.CriaNovoAluno)
// 	req,
// }
