package services

import (
	"golangBootcamp/m/common"
	"golangBootcamp/m/models"
)

func FindAllPokemons() ([]models.Pokemon, error) {
	csvPokemons, error := common.ReadCsvFile("./pokemon.csv")
	if error != nil {
		return nil, error
	}
	pokemons := []models.Pokemon{}
	for _, csvPokemon := range csvPokemons {
		pokemons = append(pokemons, models.Pokemon{Id: csvPokemon[0], Name: csvPokemon[1]})
	}
	return pokemons, nil
}

func FindAllPokemonsById(id string) (*models.Pokemon, error) {
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
