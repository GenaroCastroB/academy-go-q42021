package common

import (
	"encoding/csv"
	"fmt"
	"golangBootcamp/m/models"
	"io"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type PokemonRepo struct{}

func NewPokemonRepo() PokemonRepo {
	return PokemonRepo{}
}

func ReadCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
		return nil, err
	}
	defer file.Close()

	records, err := readFile(file)
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
		return nil, err
	}

	return records, nil
}

func readFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, err
}

func (pr PokemonRepo) GetPokemonsFromCSV() ([]models.Pokemon, error) {
	csvPokemons, error := ReadCsvFile(viper.GetString("data.pokemon.file"))
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
