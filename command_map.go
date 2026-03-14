package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pderyuga/pokedex-go/internal/pokecache"
)

type Config struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(param string, config *Config, cache *pokecache.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.Next != nil {
		url = *config.Next
	}
	item, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(item, config)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, config)
		if err != nil {
			return err
		}

		cache.Add(url, body)
	}

	for _, location := range config.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(param string, config *Config, cache *pokecache.Cache) error {
	url := config.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	item, ok := cache.Get(*url)
	if ok {
		err := json.Unmarshal(item, config)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(*url)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, config)
		if err != nil {
			return err
		}

		cache.Add(*url, body)
	}

	for _, location := range config.Results {
		fmt.Println(location.Name)
	}

	return nil
}
