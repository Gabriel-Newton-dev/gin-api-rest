package main

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/controllers"
	"github.com/Gabriel-Newton-dev/gin-api-rest/database"
	"github.com/Gabriel-Newton-dev/gin-api-rest/routes"
)

func main() {
	controllers.CallViper()
	database.ConnectDataBase()
	routes.HandleRequests()
}
