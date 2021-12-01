package requests

import (
	"encoding/json"
	"golangBootcamp/m/models"
	"io/ioutil"
	"net/http"
)

func ParsePokemonsFromApi(pokemonResponse *http.Response) ([]models.Pokemon, error) {
	responseData, err := ioutil.ReadAll(pokemonResponse.Body)
	if err != nil {
		return nil, err
	}
	var responseObject models.PokemonResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject.Pokemons, nil
}
