package routes

import (
	"pokemon-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	pokemonRoutes := router.Group("/pokemons")
	{
		pokemonRoutes.POST("/", controllers.CreatePokemon)
		pokemonRoutes.GET("/", controllers.ListPokemons)
		pokemonRoutes.PATCH("/:id", controllers.UpdatePokemon)
		pokemonRoutes.DELETE("/:id", controllers.DeletePokemon)
	}

	return router
}
