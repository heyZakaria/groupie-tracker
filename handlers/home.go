package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/error.html")
	if err != nil {
		http.Error(w, "Error page not found: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		renderErrorPage(w, http.StatusNotFound, "Page Not Found")
		return
	}

	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", response.Status)
	}

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
