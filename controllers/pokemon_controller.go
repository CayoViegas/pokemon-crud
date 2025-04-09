package controllers

import (
	"errors"
	"net/http"

	"pokemon-crud/database"
	"pokemon-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TODO: Adicionar mensagens de erro personalizadas;
// no método UpdatePokemon, nenhum dos campos é obrigatório,
// mas deve, obrigatoriamente, ser possível atualizar um, outro ou os dois campos;

// CreatePokemon godoc
// @Summary      Adiciona um novo Pokémon
// @Description  Cria um novo Pokémon com nome e nível de evolução
// @Tags         Pokémons
// @Accept       json
// @Produce      json
// @Param        pokemon  body      models.PokemonDTO  true  "Dados do Pokémon"
// @Success      201      {object}  models.Pokemon
// @Failure      400      {object}  models.ErrorResponse
// @Router       /pokemons [post]
func CreatePokemon(c *gin.Context) {
	var dto models.PokemonDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		handleValidationErrors(c, err)
		return
	}

	pokemon := models.Pokemon{
		Name:      dto.Name,
		EvolvesAt: dto.EvolvesAt,
	}

	if err := database.DB.Create(&pokemon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Errors: []string{"Error creating Pokémon in database."},
		})
		return
	}

	c.JSON(http.StatusCreated, pokemon)
}

// ListPokemons godoc
// @Summary      Lista Pokémons
// @Description  Retorna todos os Pokémons ordenados por nível de evolução, do menor para o maior
// @Tags         Pokémons
// @Produce      json
// @Success      200  {array}  models.Pokemon
// @Failure      500  {object}  models.ErrorResponse
// @Router       /pokemons [get]
func ListPokemons(c *gin.Context) {
	var pokemons []models.Pokemon

	if err := database.DB.Order("evolves_at asc").Find(&pokemons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao listar pokémons"})
		return
	}

	c.JSON(http.StatusOK, pokemons)
}

// UpdatePokemon godoc
// @Summary      Atualiza um Pokémon
// @Description  Atualiza as informações de um Pokémon pelo ID. Pelo menos um dos campos deve ser fornecido
// @Tags         Pokémons
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "ID do Pokémon"
// @Param        pokemon  body      models.PokemonUpdateDTO  true  "Dados atualizados do Pokémon (Nome e/ou Nível de Evolução)"
// @Success      200      {object}  models.Pokemon
// @Failure      400      {object}  models.ErrorResponse "Erro de validação: campos inválidos ou nenhum campo fornecido"
// @Failure      404      {object}  models.ErrorResponse "Pokémon não encontrado"
// @Router       /pokemons/{id} [patch]
func UpdatePokemon(c *gin.Context) {
	id := c.Param("id")
	var pokemon models.Pokemon

	if err := database.DB.First(&pokemon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Errors: []string{"Pokemon not found"},
		})
		return
	}

	var dto models.PokemonUpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		handleValidationErrors(c, err)
		return
	}

	if dto.Name != "" {
		pokemon.Name = dto.Name
	}

	if dto.EvolvesAt != 0 {
		pokemon.EvolvesAt = dto.EvolvesAt
	}

	if err := database.DB.Save(&pokemon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Errors: []string{"Error updating Pokémon in database."},
		})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

// DeletePokemon godoc
// @Summary      Deleta um Pokémon
// @Description  Remove um Pokémon pelo ID
// @Tags         Pokémons
// @Produce      json
// @Param        id   path      int  true  "ID do Pokémon"
// @Success      200  {string}  string  "Pokémon deletado com sucesso"
// @Failure      404  {object}  models.ErrorResponse
// @Router       /pokemons/{id} [delete]
func DeletePokemon(c *gin.Context) {
	id := c.Param("id")

	var pokemon models.Pokemon
	if err := database.DB.First(&pokemon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	database.DB.Delete(&pokemon)
	c.JSON(http.StatusOK, gin.H{"message": "Pokemon deleted"})
}

func handleValidationErrors(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = models.GetValidationMessage(fe)
		}
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: out})
		return
	}
	c.JSON(http.StatusBadRequest, models.ErrorResponse{
		Errors: []string{"Invalid request format."},
	})
}
