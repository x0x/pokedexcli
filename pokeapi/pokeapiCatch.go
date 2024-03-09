package pokeapi

import (
	"encoding/json"
	"io"
)

func CatchPokemon(pokeapiClient *Client, pokemonName *string) (CatchResponse, error) {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	url := baseUrl + *pokemonName
	var response CatchResponse
	res, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return CatchResponse{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CatchResponse{}, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return CatchResponse{}, err
	}
	return response, nil
}
