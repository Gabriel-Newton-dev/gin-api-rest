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
	StringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(StringDeConexao)) //&gorm.Config{}
	if err != nil {
		log.Println("Não foi possível conectar com banco de dados", err)
	}

	DB.AutoMigrate(&models.Aluno{})

	// para eu criar uma tabela no banco de dados usando o GORM, eu utilizo o DB.AutoMigrate
	// passando o endereço de memória da Struct que eu quero criar com uma instancia dela{}

}
