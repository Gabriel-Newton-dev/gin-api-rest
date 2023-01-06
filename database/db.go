package database

import (
	"log"

	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "user=postgres dbname=postgres password= host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}

// para eu criar uma tabela no banco de dados usando o GORM, eu utilizo o DB.AutoMigrate
// passando o endereço de memória da Struct que eu quero criar com uma instancia dela{}
