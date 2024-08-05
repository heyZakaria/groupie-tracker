package Music

import (
	"encoding/json"
	"log"
	"net/http"
)

type data struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

var MyData data

func GetApi() {
	url := "https://groupietrackers.herokuapp.com/api"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(response.Body).Decode(&MyData)
}
