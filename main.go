package main

import (
	"pokemon-crud/database"
	_ "pokemon-crud/docs"
	"pokemon-crud/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Pokemon CRUD API
// @version 1.0
// @description API para gerenciar a evolução dos pokemons
// @host localhost:8080
// @BasePath /
func main() {
	database.ConnectDatabase()
	router := routes.SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
