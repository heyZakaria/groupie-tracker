package main

import (
	"fmt"
	"net/http"

	Music "groupietracker/packages/funcs"
)

func main() {
	Music.GetApi()
	

	http.HandleFunc("/",  Music.GetArtistes)

	fmt.Println("http://localhost:9890")
	http.ListenAndServe(":9890", nil)
}
