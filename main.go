package main

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/routes"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".envn")
	viper.ReadInConfig()
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
