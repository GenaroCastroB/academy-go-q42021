package clients

import (
	"golangBootcamp/m/common/requests"
	"golangBootcamp/m/models"
	"net/http"
)

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type PokemonClient struct {
	client httpClient
	apiUrl string
}

func NewPokemonClient(apiUrl string, client httpClient) PokemonClient {
	return PokemonClient{client, apiUrl}
}

func (pkc PokemonClient) GetPokemons() ([]models.Pokemon, error) {
	response, err := pkc.client.Get(pkc.apiUrl)
	if err != nil {
		return nil, err
	}

	return requests.ParsePokemonsFromApi(response)
}
