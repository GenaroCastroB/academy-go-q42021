package controllers

import (
	"golangBootcamp/m/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPokemons(c *gin.Context) {
	pokemons, error := services.FindAllPokemons()
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pokemons})
}

func FindPokemonById(c *gin.Context) {
	id := c.Param("id")
	pokemon, error := services.FindPokemonById(id)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error})
		return
	}
	if pokemon != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": pokemon})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "pokemon not found"})
}
