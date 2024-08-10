package main

import (
	"fmt"
	"net/http"

	Music "groupietracker/packages/funcs"
)

const (
	url = "https://groupietrackers.herokuapp.com/api"
)

func main() {
	fmt.Println("http://localhost:9890")

	Music.GetApi(url)

	http.HandleFunc("/", Music.GetArtists)
	http.HandleFunc("/artist", Music.GetArtist)
	fmt.Println(Music.Location)
	
	http.ListenAndServe(":9890", nil)
}
