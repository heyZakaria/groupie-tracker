package Music

import (
	"encoding/json"
	"log"
	"net/http"
)

type urls struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

var MyApi urls

func GetApi(url string) {
	
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(response.Body).Decode(&MyApi)
	defer response.Body.Close()
	
}
