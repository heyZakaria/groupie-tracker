package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetApi() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println(string(responseData))

	json.Unmarshal(responseData, &MyData)

	// Artists-------------------------------

	ArtisteResp, err := http.Get(MyData.Artists)
	if err != nil {
		log.Fatalln(err)
	}
	ArtisteData, err := ioutil.ReadAll(ArtisteResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ArtisteData))

	// Locations-------------------------------

	LocationsResp, err := http.Get(MyData.Locations)
	if err != nil {
		log.Fatalln(err)
	}
	LocationsData, err := ioutil.ReadAll(LocationsResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(LocationsData))

	// Dates-----------------------------------

	DatesResp, err := http.Get(MyData.Dates)
	if err != nil {
		log.Fatalln(err)
	}
	DatesData, err := ioutil.ReadAll(DatesResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(DatesData))


}

type Data struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

var MyData Data

func main() {
	GetApi()
	
}
