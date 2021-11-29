package repositories

import (
	"errors"
	"golangBootcamp/m/models"
	"reflect"
	"testing"
)

type MockedPokemonCsv struct {
	expectedCSVData [][]string
	expectedError   error
}

func (mpr MockedPokemonCsv) ReadCsvFile(filePath string) ([][]string, error) {
	return mpr.expectedCSVData, mpr.expectedError
}

func TestGetPokemonsFromCSV(t *testing.T) {
	subtests := []struct {
		name         string
		mpr          PokemonCSV
		pokemonId    int
		expectedData []models.Pokemon
		expectingErr bool
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
			expectingErr: false,
		},
		{
			name: "Error on read csv file",
			mpr: &MockedPokemonCsv{
				expectedCSVData: nil,
				expectedError:   errors.New("Error"),
			},
			pokemonId:    1,
			expectedData: nil,
			expectingErr: true,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			repo := NewPokemonRepo(subtest.mpr)
			pokemons, error := repo.GetPokemonsFromCSV()
			if !reflect.DeepEqual(pokemons, subtest.expectedData) {
				t.Errorf("Expected (%v), got (%v)", subtest.expectedData, pokemons)
			}
			errExist := error != nil
			if subtest.expectingErr != errExist {
				t.Errorf("Expected (%v) error, got (%v) error", subtest.expectingErr, errExist)
			}
		})
	}
}
