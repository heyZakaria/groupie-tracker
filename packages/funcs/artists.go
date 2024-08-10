package Music

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Bad Request ", http.StatusBadRequest)

		return

	}
	if r.URL.Path != "/" {
		fmt.Fprint(w, "Page Not Found ", http.StatusNotFound)
		return
	}
	artistUrl := MyApi.Artists
	artisteRes, err := http.Get(artistUrl)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(artisteRes.Body).Decode(&Artist)
	defer artisteRes.Body.Close()

	tmp, err := template.ParseFiles("packages/pages/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmp.Execute(w, Artist)
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	// artistlocation := MyApi.Locations

	// artisteRes, err := http.Get(artistlocation)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println(artisteRes.Body)
	
	
	// json.NewDecoder(artisteRes.Body).Decode(&Location)
	// fmt.Println(artisteRes.Body)
	// defer artisteRes.Body.Close()

	r.ParseForm()
	s:= r.URL.Query().Get("id")
	fmt.Println(s)

	tmp, err := template.ParseFiles("packages/pages/artist.html")
	if err != nil {
		log.Fatal(err)
	}

	tmp.Execute(w, Location)

	
}
