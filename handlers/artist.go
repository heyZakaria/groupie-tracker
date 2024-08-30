package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func ArtistDetails(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("views/error.html")
	if err != nil {
		http.Error(w, "Error page not found: Internal Server Error", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Query().Get("Id")
	if idStr == "" {
		renderErrorPage(w, http.StatusBadRequest, "Missing artist ID")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		renderErrorPage(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	// Fetch the list of artists from the API
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

	var selectedArtist Artist
	found := false
	for _, artist := range artists {
		if artist.Id == id {
			selectedArtist = artist
			found = true
			break
		}
	}

	if !found {
		renderErrorPage(w, http.StatusNotFound, "Artist not found")
		return
	}

	// Fetch relations data
	relationURL := "https://groupietrackers.herokuapp.com/api/relation"
	relationResponse, err := http.Get(relationURL)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	defer relationResponse.Body.Close()

	relationBody, err := io.ReadAll(relationResponse.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	if relationResponse.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", relationResponse.Status)
	}

	var relations Relations
	err = json.Unmarshal(relationBody, &relations)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Find relations for the selected artist
	var artistRelations map[string][]string
	for _, relation := range relations.Index {
		if relation.Id == id {
			artistRelations = relation.DatesLocations
			break
		}
	}

	if artistRelations == nil {
		artistRelations = make(map[string][]string)
	}

	// Combine artist and relations data
	data := struct {
		Artist    Artist
		Relations map[string][]string
	}{
		Artist:    selectedArtist,
		Relations: artistRelations,
	}

	// Parse and execute the template with combined data
	tpl, err = template.ParseFiles("views/details.html")
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
