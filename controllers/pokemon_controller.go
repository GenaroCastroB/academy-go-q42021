package controllers

import (
	"fmt"
	"golangBootcamp/m/models"
	"net/http"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

type pokemonService interface {
	FindAllPokemons() ([]models.Pokemon, error)
	FindPokemonById(id int) (*models.Pokemon, error)
	FindPokemonByType(idType string, items int, itemsPerWorker int) ([]models.Pokemon, error)
	LoadPokemons() error
}

type PokemonServiceHandler struct {
	pokemonService pokemonService
}

func NewPokemonServiceHandler(pokemonService pokemonService) PokemonServiceHandler {
	return PokemonServiceHandler{pokemonService}
}

var AllowedIdTypes = map[string]bool{
	"odd":  true,
	"even": true,
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

func (pks PokemonServiceHandler) FindPokemonsConcurrently(c *gin.Context) {
	idType := c.Query("type")
	items, itemsErr := strconv.Atoi(c.Query("items"))
	itemsPerWorkers, itemsPWErr := strconv.Atoi(c.Query("items_per_workers"))
	if itemsErr != nil {
		fmt.Println("Error ", itemsErr)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid items param"})
		return
	}
	if itemsPWErr != nil {
		fmt.Println("Error ", itemsPWErr)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid items per worker param"})
		return
	}
	if _, ok := AllowedIdTypes[idType]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid type param"})
		return
	}
	start := time.Now()
	pokemons, err := pks.pokemonService.FindPokemonByType(idType, items, itemsPerWorkers)
	if err != nil {
		fmt.Println("Error ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}
	fmt.Printf("\n%2fs", time.Since(start).Seconds())
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
