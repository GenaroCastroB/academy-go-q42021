package services

import (
	"fmt"
	"golangBootcamp/m/common/requests"
	"golangBootcamp/m/models"
	"net/http"

	"github.com/spf13/viper"
)

type pokemonRepo interface {
	GetPokemonsFromCSV() ([]models.Pokemon, error)
	WritePokemonCsvFile(pokemons []models.Pokemon) (bool, error)
}

type PokemonService struct {
	repo pokemonRepo
}

func NewPokemonService(repo pokemonRepo) PokemonService {
	return PokemonService{repo}
}

func (pks PokemonService) FindAllPokemons() ([]models.Pokemon, error) {
	pokemons, error := pks.repo.GetPokemonsFromCSV()
	if error != nil {
		return nil, error
	}
	return pokemons, nil
}

func (pks PokemonService) FindPokemonById(id int) (*models.Pokemon, error) {
	pokemons, error := pks.FindAllPokemons()
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

func (pks PokemonService) LoadPokemons() (bool, error) {
	response, error := http.Get(viper.GetString("api.pokemon.url"))
	if error != nil {
		fmt.Println("Error getting data from api: ", error)
		return false, error
	}
	parsedPokemons, error := requests.ParsePokemonsFromApi(response)
	if error != nil {
		fmt.Println("Error parsing pokemon data: ", error)
		return false, error
	}

	return pks.repo.WritePokemonCsvFile(parsedPokemons)
}
