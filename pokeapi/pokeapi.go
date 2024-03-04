package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"-"`
}

var baseUrl string = "https://pokeapi.co/api/v2/location"

func GetLocations(pageUrl *string) (Response, error) {
	url := baseUrl
	if pageUrl != nil {
		url = *pageUrl
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return Response{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		return Response{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}
