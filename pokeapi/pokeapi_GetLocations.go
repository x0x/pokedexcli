package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type GetResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"-"`
	} `json:"results"`
}

var baseUrl string = "https://pokeapi.co/api/v2/location"

func GetLocations(pokeapiClient *Client, pageUrl *string) (GetResponse, error) {
	url := baseUrl
	if pageUrl != nil {
		url = *pageUrl
	}

	// cache check
	data, ok := pokeapiClient.cache.Get(url)
	var response GetResponse
	if ok {
		fmt.Println("[Cache hit]")
		response := GetResponse{}
		if err := json.Unmarshal(data, &response); err != nil {
			return GetResponse{}, err
		}
		return response, nil
	}

	fmt.Println("[Cache miss]")
	res, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return GetResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	pokeapiClient.cache.Add(url, body)
	if err := json.Unmarshal(body, &response); err != nil {
		return GetResponse{}, err
	}
	return response, nil
}
