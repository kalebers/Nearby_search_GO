package main

// precisa-se usar a latitude e longitude do lugar para pegar recomendacoes nearby: -25.440660614399487, -49.27741271669351

// to do: how to save the json request in a file?

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-25.440660614399487,-49.277412716693516&radius=10000&type=restaurant&keyword=rankby&key="
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	//os.WriteFile("go/go_exercises"+".txt", []byte(url), 0644)
}
