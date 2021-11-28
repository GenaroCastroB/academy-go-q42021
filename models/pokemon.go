package models

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PokemonResponse struct {
	Count    int       `json:"count"`
	Pokemons []Pokemon `json:"results"`
}
