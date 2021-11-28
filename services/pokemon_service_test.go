package services

import (
	"errors"
	"golangBootcamp/m/models"
	"reflect"
	"testing"
)

type MockedPokemonRepo struct {
	expectedPokemons []models.Pokemon
	expectedError    error
}

func (mpr MockedPokemonRepo) GetPokemonsFromCSV() ([]models.Pokemon, error) {
	return mpr.expectedPokemons, mpr.expectedError
}

func (mpr MockedPokemonRepo) WritePokemonCsvFile(pokemons []models.Pokemon) (bool, error) {
	return true, nil
}
func TestFindPokemonById(t *testing.T) {
	subtests := []struct {
		name         string
		mpr          pokemonRepo
		pokemonId    int
		expectedData *models.Pokemon
		expectingErr bool
	}{
		{
			name: "Happy path",
			mpr: &MockedPokemonRepo{
				expectedPokemons: []models.Pokemon{
					{Id: 1, Name: "name"},
					{Id: 2, Name: "name"},
				},
				expectedError: nil,
			},
			pokemonId:    1,
			expectedData: &models.Pokemon{Id: 1, Name: "name"},
			expectingErr: false,
		},
		{
			name: "Pokemon not found",
			mpr: &MockedPokemonRepo{
				expectedPokemons: []models.Pokemon{
					{Id: 1, Name: "name"},
					{Id: 2, Name: "name"},
				},
				expectedError: nil,
			},
			pokemonId:    3,
			expectedData: nil,
			expectingErr: false,
		},
		{
			name: "Find pokemons error",
			mpr: &MockedPokemonRepo{
				expectedPokemons: nil,
				expectedError:    errors.New("Error"),
			},
			pokemonId:    1,
			expectedData: nil,
			expectingErr: true,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			service := NewPokemonService(subtest.mpr)
			pokemon, error := service.FindPokemonById(subtest.pokemonId)
			if !reflect.DeepEqual(pokemon, subtest.expectedData) {
				t.Errorf("Expected (%v), got (%v)", subtest.expectedData, pokemon)
			}
			errExist := error != nil
			if subtest.expectingErr != errExist {
				t.Errorf("Expected (%v) error, got (%v) error", subtest.expectingErr, errExist)
			}
		})
	}
}
