package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func ListLocations(pokeapiClient *Client, location string) (ListResponse, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	url := baseUrl + location

	if data, ok := pokeapiClient.cache.Get(url); ok {
		fmt.Println("[Cache hit]")
		response := ListResponse{}
		if err := json.Unmarshal(data, &response); err != nil {
			return ListResponse{}, err
		}
		return response, nil
	}

	res, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return ListResponse{}, err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()
	if err != nil {
		return ListResponse{}, err
	}

	pokeapiClient.cache.Add(url, body)

	var response ListResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return ListResponse{}, err
	}

	return response, nil
}
