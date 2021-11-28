package requests

import (
	"encoding/json"
	"golangBootcamp/m/models"
	"io/ioutil"
	"net/http"
)

func ParsePokemonsFromApi(pokemonResponse *http.Response) ([]models.Pokemon, error) {
	responseData, error := ioutil.ReadAll(pokemonResponse.Body)
	if error != nil {
		return nil, error
	}
	var responseObject models.PokemonResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject.Pokemons, nil
}
