package main

import (
	"bytes"
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
	aluno := models.Aluno{Nome: "Student Test", CPF: "12345678910", RG: "123456789"}
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
	database.ConnectDataBase()
	r := RouterSetup()
	r.GET("/alunos", controllers.DisplaysAllStudent)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder() // para armanezar todas as informacoes do corpo da nossa resposta
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code) // assertação para verificar se de fato estamos recebendo o status code que a gente espera.
	// fmt.Println(response.Body)
	// t- teste, depois valor esperado - http.StatusOK, e o que vou receber é o status code da resposta
}

func TestSearchByCPF(t *testing.T) {
	controllers.CallViper()
	database.ConnectDataBase()
	r := RouterSetup()
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/20092060720", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudentByID(t *testing.T) {
	controllers.CallViper()
	database.ConnectDataBase()
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
	assert.Equal(t, "Student Test", StudentMock.Nome) //Firts test t, after expected value is value that we will receive in alunoMock.Name
}
func TestDeleteStudent(t *testing.T) {
	controllers.CallViper()
	database.ConnectDataBase()
	CreateStudentMock()
	r := RouterSetup()
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	searchPath := "/alunos/" + strconv.Itoa(ID)
	// searchPath := "/alunos/36" - podemos passar direto aqui o id para ele realizar a exclusão.
	req, _ := http.NewRequest("DELETE", searchPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

// r.PATCH("/alunos/:id", controllers.EditStudent)

func TestEditStudantHandle(t *testing.T) {
	controllers.CallViper()
	database.ConnectDataBase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouterSetup()
	r.PATCH("/alunos/:id", controllers.EditStudent)
	aluno := models.Aluno{Nome: "Student Test", CPF: "25696969121", RG: "987654321"}
	valueJson, _ := json.Marshal(aluno)
	PathEdit := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", PathEdit, bytes.NewBuffer(valueJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMockUpdated models.Aluno
	json.Unmarshal(response.Body.Bytes(), &studentMockUpdated)
	assert.Equal(t, "Student Test", studentMockUpdated.Nome)
	assert.Equal(t, "25696969121", studentMockUpdated.CPF)
	assert.Equal(t, "987654321", studentMockUpdated.RG)

}
