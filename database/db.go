package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaComBancoDeDados() {
	StringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(StringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Não foi possível conectar com banco de dados")
	}

	DB
	defer DB.close
}

