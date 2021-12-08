package repositories

import (
	"context"
	"encoding/csv"
	"fmt"
	"golangBootcamp/m/common"
	"golangBootcamp/m/models"
	"io"
	"os"
	"strconv"
	"sync"
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

func (pr PokemonRepo) GetPokemonsFromCSVConcurrently(idType string, items int, itemsPerWorker int) ([]models.Pokemon, error) {
	numberOfWorkers := items / itemsPerWorker
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	defer cancel()
	row := make(chan []string)
	pokemon := make(chan models.Pokemon, numberOfWorkers)
	readedItems := 0

	csvfile, err := os.Open(pr.dataFilePath)
	if err != nil {
		return nil, err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				return
			}
			row <- record
		}
		close(row)
	}()

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pokemonObtainer(ctx, pokemon, row, idType)
		}()
	}

	go func() {
		wg.Wait()
		close(pokemon)
	}()

	pokemonList := []models.Pokemon{}
	for res := range pokemon {
		if readedItems == items {
			break
		}
		pokemonList = append(pokemonList, res)
		readedItems++
	}

	return pokemonList, nil
}

func pokemonObtainer(ctx context.Context, pokemon chan models.Pokemon, src chan []string, idType string) {
	for {
		select {
		case row, ok := <-src:
			if !ok {
				return
			}
			intId, err := strconv.Atoi(row[0])
			if err != nil {
				return
			}
			switch idType {
			case "even":
				if common.Even(intId) {
					pokemon <- models.Pokemon{Id: intId, Name: row[1]}
				}
			case "odd":
				if !common.Even(intId) {
					pokemon <- models.Pokemon{Id: intId, Name: row[1]}
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
