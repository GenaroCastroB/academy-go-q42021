package controllers

import (
	"errors"
	"golangBootcamp/m/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockedPokemonServiceHandler struct {
	expectedPokemons []models.Pokemon
	expectedPokemon  *models.Pokemon
	expectedError    error
}

func (mpsh MockedPokemonServiceHandler) FindAllPokemons() ([]models.Pokemon, error) {
	return mpsh.expectedPokemons, mpsh.expectedError
}

func (mpsh MockedPokemonServiceHandler) FindPokemonById(id int) (*models.Pokemon, error) {
	return mpsh.expectedPokemon, mpsh.expectedError
}

func (mpsh MockedPokemonServiceHandler) LoadPokemons() error {
	return mpsh.expectedError
}

func (mpsh MockedPokemonServiceHandler) FindPokemonByType(idType string, items int, itemsPerWorker int) ([]models.Pokemon, error) {
	return nil, nil
}

func TestFindPokemons(t *testing.T) {
	subtests := []struct {
		name           string
		mpsh           pokemonService
		expectedStatus int
	}{
		{
			name:           "Happy path",
			mpsh:           &MockedPokemonServiceHandler{},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Error on read csv file",
			mpsh: &MockedPokemonServiceHandler{
				expectedError: errors.New("Error"),
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			pks := NewPokemonServiceHandler(subtest.mpsh)
			pks.FindPokemons(c)
			assert.Equal(t, w.Code, subtest.expectedStatus)
		})
	}
}

func TestFindPokemonById(t *testing.T) {
	subtests := []struct {
		name           string
		mpsh           pokemonService
		pokemonId      string
		expectedStatus int
	}{
		{
			name: "Pokemon found",
			mpsh: &MockedPokemonServiceHandler{
				expectedPokemon: &models.Pokemon{Id: 1, Name: "name"},
			},
			pokemonId:      "1",
			expectedStatus: http.StatusOK,
		},
		{
			name: "Pokemon not fond",
			mpsh: &MockedPokemonServiceHandler{
				expectedPokemon: nil,
			},
			pokemonId:      "1",
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "Bad request",
			mpsh: &MockedPokemonServiceHandler{
				expectedError: errors.New("Error"),
			},
			pokemonId:      "a",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Error finding pokemon",
			mpsh: &MockedPokemonServiceHandler{
				expectedError: errors.New("Error"),
			},
			pokemonId:      "1",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			pks := NewPokemonServiceHandler(subtest.mpsh)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: subtest.pokemonId})
			pks.FindPokemonById(c)
			assert.Equal(t, w.Code, subtest.expectedStatus)
		})
	}
}

func TestLoadPokemons(t *testing.T) {
	subtests := []struct {
		name           string
		mpsh           pokemonService
		expectedStatus int
	}{
		{
			name:           "Happy path",
			mpsh:           &MockedPokemonServiceHandler{},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Error on loading pokemons",
			mpsh: &MockedPokemonServiceHandler{
				expectedError: errors.New("Error"),
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			pks := NewPokemonServiceHandler(subtest.mpsh)
			pks.LoadPokemons(c)
			assert.Equal(t, w.Code, subtest.expectedStatus)
		})
	}
}
