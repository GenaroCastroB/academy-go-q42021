package services

import (
	"fmt"
	"golangBootcamp/m/models"
)

type pokemonRepo interface {
	GetPokemonsFromCSV() ([]models.Pokemon, error)
	WritePokemonCsvFile(pokemons []models.Pokemon) error
}

type pokemonClient interface {
	GetPokemons() ([]models.Pokemon, error)
}

type PokemonService struct {
	repo   pokemonRepo
	client pokemonClient
}

func NewPokemonService(repo pokemonRepo, client pokemonClient) PokemonService {
	return PokemonService{repo, client}
}

func (pks PokemonService) FindAllPokemons() ([]models.Pokemon, error) {
	pokemons, err := pks.repo.GetPokemonsFromCSV()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (pks PokemonService) FindPokemonById(id int) (*models.Pokemon, error) {
	pokemons, err := pks.FindAllPokemons()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return &pokemon, nil
		}
	}
	return nil, nil
}

func (pks PokemonService) LoadPokemons() error {
	parsedPokemons, err := pks.client.GetPokemons()
	if err != nil {
		fmt.Println("Error getting data from api: ", err)
		return err
	}

	return pks.repo.WritePokemonCsvFile(parsedPokemons)
}
