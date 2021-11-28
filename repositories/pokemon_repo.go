package repositories

import (
	"encoding/csv"
	"fmt"
	"golangBootcamp/m/models"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type PokemonCSV interface {
	ReadCsvFile(filePath string) ([][]string, error)
}

type PokemonRepo struct {
	pokemonCSV PokemonCSV
}

func NewPokemonRepo(pokemonCsv PokemonCSV) PokemonRepo {
	return PokemonRepo{pokemonCsv}
}

func (pr PokemonRepo) GetPokemonsFromCSV() ([]models.Pokemon, error) {
	fmt.Println("FilePath", viper.GetString("data.pokemon.file"))
	csvPokemons, error := pr.pokemonCSV.ReadCsvFile(viper.GetString("data.pokemon.file"))
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

func (pr PokemonRepo) WritePokemonCsvFile(pokemons []models.Pokemon) (bool, error) {
	file, error := os.Create(viper.GetString("data.pokemon.file"))
	if error != nil {
		fmt.Println("Cannot create file", error)
		return false, error
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i, pokemon := range pokemons {
		csvRow := []string{strconv.Itoa(i), pokemon.Name}
		error := writer.Write(csvRow)
		if error != nil {
			fmt.Println("Cannot write to file", error)
			return false, error
		}
	}
	return true, nil
}
