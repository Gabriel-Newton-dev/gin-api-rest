package database

import (
	"fmt"
	"log"

	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	DB_NAME := viper.Get("DB_NAME")
	DB_PASSWORD := viper.Get("DB_PASSWORD")
	DB_USER := viper.Get("DB_USER")

	stringDeConexao := fmt.Sprintf("user=%s dbname=%s password=%s host=localhost sslmode=disable", DB_USER, DB_NAME, DB_PASSWORD)
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}

// para eu criar uma tabela no banco de dados usando o GORM, eu utilizo o DB.AutoMigrate
// passando o endereço de memória da Struct que eu quero criar com uma instancia dela{}
