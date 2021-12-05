package services

import (
	"errors"
	"golangBootcamp/m/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockedPokemonRepo struct {
	expectedPokemons []models.Pokemon
	expectedError    error
}

func (mpr MockedPokemonRepo) GetPokemonsFromCSV() ([]models.Pokemon, error) {
	return mpr.expectedPokemons, mpr.expectedError
}

func (mpr MockedPokemonRepo) WritePokemonCsvFile(pokemons []models.Pokemon) error {
	return nil
}

func (mpr MockedPokemonRepo) GetPokemonsFromCSVConcurrently(idType string, items int, itemsPerWorker int) ([]models.Pokemon, error) {
	return nil, nil
}

type MockedPokemonClient struct {
	expectedPokemons []models.Pokemon
	expectedError    error
}

func (mpc MockedPokemonClient) GetPokemons() ([]models.Pokemon, error) {
	return mpc.expectedPokemons, mpc.expectedError
}

func TestFindPokemonById(t *testing.T) {
	subtests := []struct {
		name         string
		mpr          pokemonRepo
		mpc          pokemonClient
		pokemonId    int
		expectedData *models.Pokemon
		expectedErr  error
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
			mpc:          nil,
			pokemonId:    1,
			expectedData: &models.Pokemon{Id: 1, Name: "name"},
			expectedErr:  nil,
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
			mpc:          nil,
			pokemonId:    3,
			expectedData: nil,
			expectedErr:  nil,
		},
		{
			name: "Find pokemons error",
			mpr: &MockedPokemonRepo{
				expectedPokemons: nil,
				expectedError:    errors.New("Error"),
			},
			mpc:          nil,
			pokemonId:    1,
			expectedData: nil,
			expectedErr:  errors.New("Error"),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			service := NewPokemonService(subtest.mpr, subtest.mpc)
			pokemon, err := service.FindPokemonById(subtest.pokemonId)
			assert.Equal(t, pokemon, subtest.expectedData, "they should be equal")
			assert.Equal(t, err, subtest.expectedErr)
		})
	}
}

func TestLoadPokemons(t *testing.T) {
	subtests := []struct {
		name         string
		mpr          pokemonRepo
		mpc          pokemonClient
		pokemonId    int
		expectedData *models.Pokemon
		expectedErr  error
	}{
		{
			name: "Happy path",
			mpr:  &MockedPokemonRepo{},
			mpc: &MockedPokemonClient{
				expectedError: nil,
			},
			pokemonId:    1,
			expectedData: &models.Pokemon{Id: 1, Name: "name"},
			expectedErr:  nil,
		},
		{
			name: "Error getting pokemons from api",
			mpr:  &MockedPokemonRepo{},
			mpc: &MockedPokemonClient{
				expectedError: errors.New("Error"),
			},
			expectedData: nil,
			expectedErr:  errors.New("Error"),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			service := NewPokemonService(subtest.mpr, subtest.mpc)
			err := service.LoadPokemons()
			assert.Equal(t, err, subtest.expectedErr)
		})
	}
}
