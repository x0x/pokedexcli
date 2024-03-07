package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"fmt"
	"pokedexcli/pokecache"
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

var cache pokecache.Cache = pokecache.NewCache(time.Duration(5 * time.Minute))

func GetLocations(pageUrl *string) (Response, error) {
	url := baseUrl
	if pageUrl != nil {
		url = *pageUrl
	}

	// cache check
	data, ok := cache.Get(url)
	var response Response
	if ok {
    fmt.Println("[Cache hit]")
    response :=  Response{}
		if err := json.Unmarshal(data, &response); err != nil {
			return Response{}, err
		}
  	return response, nil
	}
  
  fmt.Println("[Cache miss]")
	res, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}

	body, err := io.ReadAll(res.Body)
  if err != nil {
    return response, err
  }

	defer res.Body.Close()
  
  cache.Add(url, body)
	if err := json.Unmarshal(body, &response); err != nil {
		return Response{}, err
	}
	return response, nil
}
