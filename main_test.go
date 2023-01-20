package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
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
	if response.Code != http.StatusOK {
		t.Fatalf("Status Error: valor recebido foi %d e o esperado era %d", response.Code, http.StatusOK)
	}
}
