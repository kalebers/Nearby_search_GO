package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Config struct {
	ApiKey   string `json:"apiKey"`
	Location string `json:"location"`
	Radius   string `json:"radius"`
	Keyword  string `json:"keyword"`
}

type Response struct {
	Results []struct {
		Name   string  `json:"name"`
		Rating float64 `json:"rating"`
	} `json:"results"`
}

func main() {
	// Read the configuration file
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON into a Go data structure
	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	// Use the values from the configuration file
	apiKey := config.ApiKey
	location := config.Location
	radius := config.Radius
	keyword := config.Keyword

	// Build the URL for the API request
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s&radius=%s&keyword=%s&key=%s", location, radius, keyword, apiKey)

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

	// Sort the results by rating (highest to lowest)
	sort.Slice(data.Results, func(i, j int) bool {
		return data.Results[i].Rating > data.Results[j].Rating
	})

	// Print the names and ratings of the top 10 restaurants
	for i := 0; i < 10; i++ {
		result := data.Results[i]
		fmt.Printf("%d. %s - Rating: %.1f\n", i+1, result.Name, result.Rating)
	}

	// Save the data to a local file
	err = ioutil.WriteFile("data.json", body, 0644)
	if err != nil {
		panic(err)
	}
}
