package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var ID int

func RouterSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DeleteStudentMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)

}

func TestCheckEndpointSalutation(t *testing.T) {
	r := RouterSetup()
	r.GET("/:nome", controllers.Salutation)
	req, _ := http.NewRequest("GET", "/gabriel", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "they should be the same")
	mockDaResposta := `{"API says":"Welcome gabriel to our API."}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockDaResposta, string(responseBody))

}

func TestListingAllStudentHandler(t *testing.T) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	database.ConectaComBancoDeDados()
	r := RouterSetup()
	r.GET("/alunos", controllers.DisplayAllStudent)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder() // para armanezar todas as informacoes do corpo da nossa resposta
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code) // assertação para verificar se de fato estamos recebendo o status code que a gente espera.
	// fmt.Println(response.Body)
	// t- teste, depois valor esperado - http.StatusOK, e o que vou receber é o status code da resposta
}

func TestSearchByCPF(t *testing.T) {
	controllers.CallViper()
	database.ConectaComBancoDeDados()
	r := RouterSetup()
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/20092060720", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudentByID(t *testing.T) {
	controllers.CallViper()
	database.ConectaComBancoDeDados()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouterSetup()
	r.GET("/alunos/:id", controllers.SearchStudentbyID)
	SearchPath := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", SearchPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var StudentMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &StudentMock)
	assert.Equal(t, "Aluno Teste", StudentMock.Nome) // Firts test t, após valor esperado e valor que iremos receber no alunoMock.Name
}
