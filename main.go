package main

import (
	"fmt"
	"net/http"

	Music "groupietracker/packages/funcs"
)

const (
	api = "https://groupietrackers.herokuapp.com/api"
)

func main() {
	fmt.Println("http://localhost:3040")

	Music.GetApi(api)
	
	dir := http.FileServer(http.Dir("packages/public"))

	http.Handle("/packages/public/", http.StripPrefix("/packages/public/", dir))

	http.HandleFunc("/", Music.GetArtists)

	http.ListenAndServe(":3040", nil)
}
