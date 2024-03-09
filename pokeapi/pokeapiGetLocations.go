package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

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

	res, err := pokeapiClient.httpClient.Get(url)
	if err != nil {
		return GetResponse{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	pokeapiClient.cache.Add(url, body)
	if err := json.Unmarshal(body, &response); err != nil {
		return GetResponse{}, err
	}
	return response, nil
}
