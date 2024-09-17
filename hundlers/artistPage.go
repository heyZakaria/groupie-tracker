package Music

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func getCoordinates(location string) (float64, float64, error) {
	geoapifyURL := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", location, "84a9a08245a141e299e5a1fd45b3dbd8")
	resp, err := http.Get(geoapifyURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, 0, err
	}

	features := result["features"].([]interface{})
	if len(features) > 0 {
		geometry := features[0].(map[string]interface{})["geometry"].(map[string]interface{})
		coordinates := geometry["coordinates"].([]interface{})
		return coordinates[1].(float64), coordinates[0].(float64), nil
	}

	return 0, 0, fmt.Errorf("no coordinates found")
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	// Extract artist ID from query parameters
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		renderErrorPage(w, http.StatusBadRequest, "Artist ID is missing")
		return
	}

	// Fetch artist data from the API
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", artistID)
	resp, err := http.Get(artistURL)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Error fetching artist data")
		return
	}
	defer resp.Body.Close()

	// Decode artist data
	var artist Artist
	err = json.NewDecoder(resp.Body).Decode(&artist)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Error decoding artist data")
		return
	}

	// Fetch locations data for the artist
	locationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", artistID)
	locationResp, err := http.Get(locationURL)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Error fetching location data")
		return
	}
	defer locationResp.Body.Close()

	// Decode location data
	var location Location
	err = json.NewDecoder(locationResp.Body).Decode(&location)
	if err != nil {
		renderErrorPage(w, http.StatusNotFound, "Page Not Found")
		return
	}

	// Fetch date data for the artist
	dateURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", artistID)
	dateResp, err := http.Get(dateURL)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Error fetching date data")
		return
	}
	defer dateResp.Body.Close()

	// Decode date data
	var date Date
	err = json.NewDecoder(dateResp.Body).Decode(&date)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Error decoding date data")
		return
	}

	// Fetch coordinates for each location
	var latLngs []LatLng
	for _, loc := range location.Locations {
		lat, lng, err := getCoordinates(loc)
		if err != nil {
			renderErrorPage(w, http.StatusInternalServerError, "Error fetching coordinates")
			return
		}
		latLngs = append(latLngs, LatLng{Latitude: lat, Longitude: lng})
	}
	// Prepare the location data to pass to JavaScript
	locationData := []map[string]float64{}
	for _, latLng := range latLngs {
		locationData = append(locationData, map[string]float64{
			"lat": latLng.Latitude,
			"lng": latLng.Longitude,
		})
	}

	// Convert locationData to JSON format
	locationDataJSON, err := json.Marshal(locationData)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Error converting location data to JSON")
		return
	}

	// Fetch relation data for the artist
	relationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", artistID)
	relationResp, err := http.Get(relationURL)
	if err != nil {
		http.Error(w, "Error fetching relation data", http.StatusInternalServerError)
		return
	}
	defer relationResp.Body.Close()

	// Decode relation data
	var relation Relation
	err = json.NewDecoder(relationResp.Body).Decode(&relation)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Error decoding relation data")
		return
	}
	// Pass all the required data to the template
	pageData := ArtistPageData{
		Artist:       artist,
		Location:     location,
		LocationData: string(locationDataJSON), // Pass JSON formatted data as string
		Date:         date,
		Relation:     relation, // Add the relation data to the struct
	}

	// Parse the template
	temp, err := template.ParseFiles("views/artist.html")
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Execute the template with the pageData
	err = temp.Execute(w, pageData)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
