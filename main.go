package main

// precisa-se usar a latitude e longitude do lugar para pegar recomendacoes nearby:
// to use the maps api I have to enable the feature in gcp with the key

// to do: how to save the json request in a file? done

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiKey     = ""
	location   = ""     // latitude,longitude of location
	radius     = "5000" // radius in meters
	keyword    = "restaurant"
	maxResults = 10
)

type Response struct {
	Results []struct {
		Name   string  `json:"name"`
		Rating float64 `json:"rating"`
	} `json:"results"`
}

func main() {
	// Build the URL for the API request
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s&radius=%s&keyword=%s&maxresults=%d&key=%s", location, radius, keyword, maxResults, apiKey)

	// Make the API request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	// Unmarshal the JSON into a Go data structure
	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Print the names and ratings of the top 10 restaurants
	for i, result := range data.Results {
		fmt.Printf("%d. %s - Rating: %.1f\n", i+1, result.Name, result.Rating)
	}

	// Save the data to a local file
	err = ioutil.WriteFile("data.json", body, 0644)
	if err != nil {
		panic(err)
	}
}
