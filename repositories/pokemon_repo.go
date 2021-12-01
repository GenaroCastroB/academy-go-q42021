package repositories

import (
	"fmt"
	"golangBootcamp/m/models"
	"strconv"
)

type PokemonCSV interface {
	ReadCsvFile(filePath string) ([][]string, error)
	WriteCsvFile(filePath string, data [][]string) error
}

type PokemonRepo struct {
	pokemonCSV   PokemonCSV
	dataFilePath string
}

func NewPokemonRepo(pokemonCsv PokemonCSV, dataFilePath string) PokemonRepo {
	return PokemonRepo{pokemonCsv, dataFilePath}
}

func (pr PokemonRepo) GetPokemonsFromCSV() ([]models.Pokemon, error) {
	csvPokemons, err := pr.pokemonCSV.ReadCsvFile(pr.dataFilePath)
	if err != nil {
		return nil, err
	}
	pokemons := []models.Pokemon{}
	for _, csvPokemon := range csvPokemons {
		intId, err := strconv.Atoi(csvPokemon[0])
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, models.Pokemon{Id: intId, Name: csvPokemon[1]})
	}
	return pokemons, nil
}

func (pr PokemonRepo) WritePokemonCsvFile(pokemons []models.Pokemon) error {
	pokemonRows := [][]string{}
	for i, pokemon := range pokemons {
		pokemonRows = append(pokemonRows, []string{strconv.Itoa(i), pokemon.Name})
	}
	err := pr.pokemonCSV.WriteCsvFile(pr.dataFilePath, pokemonRows)
	if err != nil {
		fmt.Println("Error saving pokemons on file", err)
		return err
	}
	return nil
}
