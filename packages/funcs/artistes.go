package Music

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type artistedata []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var ArtisteData artistedata

func GetArtistes(w http.ResponseWriter, r *http.Request) {
	artisteUrl := MyData.Artists
	artisteRes, err := http.Get(artisteUrl)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(artisteRes.Body).Decode(&ArtisteData)

	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Bad Request ", http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		fmt.Fprint(w, "Page Not Found ", http.StatusNotFound)
		return
	}

	parsed, err := template.ParseFiles("./packages/pages/index.html")
	if err != nil {

		fmt.Fprint(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}

	for _,v := range ArtisteData{

		parsed.ExecuteTemplate(w, "index.html",v)
	}
}
