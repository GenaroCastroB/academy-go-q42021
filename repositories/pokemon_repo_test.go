package repositories

import (
	"errors"
	"golangBootcamp/m/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockedPokemonCsv struct {
	expectedCSVData [][]string
	expectedError   error
	writeCsvError   error
}

func (mpr MockedPokemonCsv) ReadCsvFile(filePath string) ([][]string, error) {
	return mpr.expectedCSVData, mpr.expectedError
}

func (mpr MockedPokemonCsv) WriteCsvFile(filePath string, data [][]string) error {
	return mpr.writeCsvError
}

func TestGetPokemonsFromCSV(t *testing.T) {
	subtests := []struct {
		name         string
		mpr          PokemonCSV
		pokemonId    int
		expectedData []models.Pokemon
		expectedErr  error
	}{
		{
			name: "Happy path",
			mpr: &MockedPokemonCsv{
				expectedCSVData: [][]string{
					{"1", "pokemon1"},
					{"2", "pokemon2"},
				},
				expectedError: nil,
			},
			pokemonId: 1,
			expectedData: []models.Pokemon{
				{Id: 1, Name: "pokemon1"},
				{Id: 2, Name: "pokemon2"},
			},
			expectedErr: nil,
		},
		{
			name: "Error on read csv file",
			mpr: &MockedPokemonCsv{
				expectedCSVData: nil,
				expectedError:   errors.New("Error"),
			},
			pokemonId:    1,
			expectedData: nil,
			expectedErr:  errors.New("Error"),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			repo := NewPokemonRepo(subtest.mpr, "")
			pokemons, err := repo.GetPokemonsFromCSV()
			assert.Equal(t, pokemons, subtest.expectedData, "they should be equal")
			assert.Equal(t, err, subtest.expectedErr)
		})
	}
}

func TestWritePokemonCsvFile(t *testing.T) {
	subtests := []struct {
		name        string
		mpr         PokemonCSV
		pokemonList []models.Pokemon
		expectedErr error
	}{
		{
			name: "Happy path",
			mpr: &MockedPokemonCsv{
				writeCsvError: nil,
			},
			pokemonList: []models.Pokemon{
				{Id: 1, Name: "name"},
			},
			expectedErr: nil,
		},
		{
			name: "Error writing csv file",
			mpr: &MockedPokemonCsv{
				writeCsvError: errors.New("Error"),
			},
			pokemonList: []models.Pokemon{
				{Id: 1, Name: "name"},
			},
			expectedErr: errors.New("Error"),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			repo := NewPokemonRepo(subtest.mpr, "")
			err := repo.WritePokemonCsvFile(subtest.pokemonList)
			assert.Equal(t, err, subtest.expectedErr)
		})
	}
}
