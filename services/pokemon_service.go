package services

import (
	"golangBootcamp/m/common"
	"golangBootcamp/m/models"
)

func FindAllPokemons() ([]models.Pokemon, error) {
	pokemons, error := common.GetPokemonsFromCSV()
	if error != nil {
		return nil, error
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
