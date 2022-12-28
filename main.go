package main

// precisa-se usar a latitude e longitude do lugar para pegar recomendacoes nearby: -25.435463328482644, -49.2767323365602

// to do: how to implement the lat and long?
// what parameters to pass to the api link?
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-25.435463328482644, -49.2767323365602&radius=1500&type=restaurant&keyword=cruise&key=YOUR_API_KEY"
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
}

/*
func main() {
	response, err := http.Get("https://maps.googleapis.com/maps/api/place/nearbysearch/outputFormat?parameters")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
*/
