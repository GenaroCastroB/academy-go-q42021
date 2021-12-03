package clients

import (
	"bytes"
	"errors"
	"golangBootcamp/m/models"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockedPokemonClient struct {
	expectedResponse *http.Response
	expectedErr      error
}

func (mpc MockedPokemonClient) Get(url string) (resp *http.Response, err error) {
	return mpc.expectedResponse, mpc.expectedErr
}

func TestFindPokemons(t *testing.T) {
	json := `{"count": 2,"results":[{"name": "pokemon1"},{"name": "pokemon2"}]}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	subtests := []struct {
		name           string
		mpc            httpClient
		expectedResult []models.Pokemon
		expectedErr    error
	}{
		{
			name: "Happy path",
			mpc: &MockedPokemonClient{
				expectedResponse: &http.Response{
					StatusCode: 200,
					Body:       r,
				},
			},
			expectedResult: []models.Pokemon{
				{Id: 0, Name: "pokemon1"},
				{Id: 0, Name: "pokemon2"},
			},
		},
		{
			name: "Error getting pokemons",
			mpc: &MockedPokemonClient{
				expectedErr: errors.New("Error"),
			},
			expectedErr: errors.New("Error"),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			mpc := NewPokemonClient("", subtest.mpc)
			pokemons, err := mpc.GetPokemons()
			assert.Equal(t, pokemons, subtest.expectedResult, "they should be equal")
			assert.Equal(t, err, subtest.expectedErr)
		})
	}
}
