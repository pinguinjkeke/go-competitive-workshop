package cmd

import (
	"encoding/json"
	"errors"
	"net/http"
)

type PokemonApiClient struct {
	Client *http.Client
}

func NewPokemonApiClient() *PokemonApiClient {
	p := new(PokemonApiClient)
	p.Client = &http.Client{}

	return p
}

func (r *PokemonApiClient) FindPokemon(name string) (p Pokemon, err error) {
	resp, err := r.Client.Get("https://pokeapi.co/api/v2/pokemon/" + name)

	if err != nil {
		return p, errors.New("Failed to receive HTTP response.")
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&p)

	if err != nil {
		return p, errors.New("Failed to parse JSON body")
	}

	return p, nil
}
