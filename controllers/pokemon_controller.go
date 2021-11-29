package controllers

import (
	"fmt"
	"golangBootcamp/m/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type pokemonService interface {
	FindAllPokemons() ([]models.Pokemon, error)
	FindPokemonById(id int) (*models.Pokemon, error)
	LoadPokemons() (bool, error)
}

type PokemonServiceHandler struct {
	pokemonService pokemonService
}

func NewPokemonServiceHandler(pokemonService pokemonService) PokemonServiceHandler {
	return PokemonServiceHandler{pokemonService}
}

func (pks PokemonServiceHandler) FindPokemons(c *gin.Context) {
	pokemons, error := pks.pokemonService.FindAllPokemons()
	if error != nil {
		fmt.Println("Error ", error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pokemons})
}

func (pks PokemonServiceHandler) FindPokemonById(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		fmt.Println("Error ", error)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id param"})
		return
	}
	pokemon, error := pks.pokemonService.FindPokemonById(id)
	if error != nil {
		fmt.Println("Error ", error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}
	if pokemon != nil {
		c.JSON(http.StatusOK, gin.H{"data": pokemon})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "pokemon not found"})
}

func (pks PokemonServiceHandler) LoadPokemons(c *gin.Context) {
	_, error := pks.pokemonService.LoadPokemons()
	if error != nil {
		fmt.Println("Error ", error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
