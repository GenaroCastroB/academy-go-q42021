package services

import (
	"golangBootcamp/m/common"
	"golangBootcamp/m/models"
	"strconv"
)

func FindAllPokemons() ([]models.Pokemon, error) {
	csvPokemons, error := common.ReadCsvFile("./pokemon.csv")
	if error != nil {
		return nil, error
	}
	pokemons := []models.Pokemon{}
	for _, csvPokemon := range csvPokemons {
		intId, error := strconv.Atoi(csvPokemon[0])
		if error != nil {
			return nil, error
		}
		pokemons = append(pokemons, models.Pokemon{Id: intId, Name: csvPokemon[1]})
	}
	return pokemons, nil
}

func FindPokemonById(id int) (*models.Pokemon, error) {
	pokemons, error := FindAllPokemons()
	if error != nil {
		return nil, error
	}
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return &pokemon, nil
		}
	}
	return nil, nil
}
