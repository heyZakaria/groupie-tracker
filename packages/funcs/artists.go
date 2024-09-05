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
	tmp, err := template.ParseFiles("packages/pages/error.html")
	if err != nil {
		http.Error(w, "The error page not found: Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Check if the artist data is already populated
	if len(Artist) == 0 {
		// Fetch the artist data from the API if not populated
		artistUrl := Urls.ArtistsUrl
		artisteRes, err := http.Get(artistUrl)
		if err != nil {
			renderErrorPage(w, http.StatusInternalServerError, "Failed to fetch artist data")
			return
		}

		defer artisteRes.Body.Close()

		err = json.NewDecoder(artisteRes.Body).Decode(&Artist)
		if err != nil {
			renderErrorPage(w, http.StatusInternalServerError, "Failed to decode artist data")
			return
		}
		var ln int
		for i, v := range Artist {
			ln = len(v.Members)
			Artist[i].Num = ln
		}
	}

	// Check if the request is for the home page or a specific artist
	if r.URL.Path == "/" {
		// Serve the homepage
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
	} else if strings.Contains(r.URL.Path, "/artist/") {
		// Handle artist-specific page
		StrId := r.URL.Path[len("/artist/"):]

		Id, err := strconv.Atoi(StrId)
		if err != nil {
			renderErrorPage(w, http.StatusNotFound, "Invalid artist ID")
			return
		}
		if Id < 1 || Id > len(Artist) {
			renderErrorPage(w, http.StatusNotFound, "Artist not found")
			return
		}

		FetchArtistData(Id, w)

	} else {
		// Handle other routes as 404
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

	for i, v := range Artist[Id].Location.Locations {

		v = strings.Replace(v, "-", " ", -1)
		v = strings.Replace(v, "_", " ", -1)

		Artist[Id].Location.Locations[i] = v

	}
	
	GeocodeLocations(Artist[Id].Location.Locations, w, Id)
	/* jsonLatLon, err := json.Marshal(Artist[Id].LongLat.Both)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
 */
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
	//fmt.Println(Lonlat.Both)

	
	//fmt.Println(jsonLatLon[len(jsonLatLon)-1])

	

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

func GeocodeLocations(l []string, w http.ResponseWriter, id int) {
	var co []float64

	for _, v := range l {
		i := 0

		v = strings.ReplaceAll(v, " ", "%20")
		GeocodeApi := "https://api.geoapify.com/v1/geocode/search?text=" + v + "&limit=1&format=json&apiKey=1ba7c108667243f0ba279f68c82e9b86"
		GeoResp, err := http.Get(GeocodeApi)
		if err != nil {
			renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		}
		defer GeoResp.Body.Close()
		json.NewDecoder(GeoResp.Body).Decode(&Lonlat)

		if len(Lonlat.Both) < len(l) {

			co = append(co, Lonlat.Results[i].Lat)
			co = append(co, Lonlat.Results[i].Lon)
			Lonlat.Both = append(Lonlat.Both, co)
			co = nil

		}

	}

	// GeoResp.Body
}
