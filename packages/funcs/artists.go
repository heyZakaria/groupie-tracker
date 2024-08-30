package Music

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type ErrorData struct {
	StatusCode int
	Message    string
}

// this function fetch data from a given API URL
func GetApi(api string) {
	response, err := http.Get(api) // send an HTTP GET reaquest to the URL=api to fetch the data from the server
	if err != nil {                // if something wrong went with the HTTP request
		log.Fatal(err) // this line for print the error and call the os.Exit(1) to stop the program
	}
	// decode the data from json type to something readable by go
	json.NewDecoder(response.Body).Decode(&Urls)
	defer response.Body.Close() // Closes Resources: avoid memory leaks
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	//Check the existance of error.html template firstly
	tmp, err := template.ParseFiles("packages/pages/error.html")
	if err != nil {
		http.Error(w, "The error page not found: Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	if r.Method != http.MethodGet {
		renderErrorPage(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if r.URL.Path == "/" {

		artistUrl := Urls.ArtistsUrl
		artisteRes, err := http.Get(artistUrl)
		if err != nil {
			log.Fatal(err)
		}

		json.NewDecoder(artisteRes.Body).Decode(&Artist)
		defer artisteRes.Body.Close()

		tmp, err = template.ParseFiles("packages/pages/index.html")
		if err != nil {
			renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		err = tmp.Execute(w, Artist)
		if err != nil {
			renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

	} else if strings.Contains(r.URL.Path, "/artists/") {

		StrId := r.URL.Path[len("/artists/"):]

		Id, err := strconv.Atoi(StrId)
		if err != nil {
			fmt.Fprint(w, "Page Not Found ", http.StatusNotFound)
			fmt.Println(err, Id)
			return
		}
		if Id < 1 || Id > 52 {
			renderErrorPage(w, http.StatusNotFound, "Page Not Found")
			return
		}

		FetchArtistData(Id, w)

		return

	} else {
		renderErrorPage(w, http.StatusNotFound, "Page Not Found")
		return
	}
}

func FetchArtistData(Id int, w http.ResponseWriter) {
	for _, v := range Artist {
		if v.ID == Id {
			Id -= 1
			Artist[Id] = v
		}
	}

	LocoURL := Artist[Id].LocationsURL

	LocoResponse, err := http.Get(LocoURL)
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(LocoResponse.Body).Decode(&Artist[Id].Location)

	defer LocoResponse.Body.Close()

	// Dates

	DatesURL := Artist[Id].ConcertDatesURL

	DateResponse, err := http.Get(DatesURL)
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(DateResponse.Body).Decode(&Artist[Id].Date)

	defer DateResponse.Body.Close()

	// Relations

	RelationURL := Artist[Id].RelationsURL

	RelationResponse, err := http.Get(RelationURL)
	if err != nil {
		fmt.Fprint(w, "Internal server error", http.StatusInternalServerError)
		log.Fatal(err)
	}
	json.NewDecoder(RelationResponse.Body).Decode(&Artist[Id].Relation)

	defer DateResponse.Body.Close()

	tmp, err := template.ParseFiles("packages/pages/artists.html")
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmp.Execute(w, Artist[Id])
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func renderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	errorData := ErrorData{
		StatusCode: statusCode,
		Message:    message,
	}
	tmpl, err := template.ParseFiles("packages/pages/error.html")
	if err != nil {
		http.Error(w, "Error loading error page", http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, errorData)
	if err != nil {
		http.Error(w, "Error rendering error page", http.StatusInternalServerError)
	}
}
