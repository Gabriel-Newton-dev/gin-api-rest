package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *gorm.DB {
	conexao := "host=localhost user=root password=root dbname=root port=5452 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conexao), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Comunição OK com Banco de Dados.")
	return db
}
