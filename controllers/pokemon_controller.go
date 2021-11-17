package controllers

import (
	"fmt"
	"golangBootcamp/m/services"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func FindPokemons(c *gin.Context) {
	pokemons, error := services.FindAllPokemons()
	if error != nil {
		fmt.Println("Error ", error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pokemons})
}

func FindPokemonById(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		fmt.Println("Error ", error)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id param"})
		return
	}
	pokemon, error := services.FindPokemonById(id)
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
