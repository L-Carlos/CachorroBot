package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// DogResponse is the response from https://dog.ceo/api/breeds/image/random
type DogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func getDog() (string, error) {
	var dog DogResponse
	var url string = "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	if err := json.NewDecoder(resp.Body).Decode(&dog); err != nil {
		log.Printf("could not decode incoming dog %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	return dog.Message, nil

}
