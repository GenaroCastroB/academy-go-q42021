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
	LoadPokemons() error
}

type PokemonServiceHandler struct {
	pokemonService pokemonService
}

func NewPokemonServiceHandler(pokemonService pokemonService) PokemonServiceHandler {
	return PokemonServiceHandler{pokemonService}
}

func (pks PokemonServiceHandler) FindPokemons(c *gin.Context) {
	pokemons, err := pks.pokemonService.FindAllPokemons()
	if err != nil {
		fmt.Println("Error ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pokemons})
}

func (pks PokemonServiceHandler) FindPokemonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id param"})
		return
	}
	pokemon, err := pks.pokemonService.FindPokemonById(id)
	if err != nil {
		fmt.Println("Error ", err)
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
	err := pks.pokemonService.LoadPokemons()
	if err != nil {
		fmt.Println("Error ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
