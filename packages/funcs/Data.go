package Music

type Artists []struct {
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

var Artist Artists

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}


var Location Locations

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

var Date Dates


type Relations struct {
	Index []struct {
		ID             int `json:"id"`
		DatesLocations struct {
			DunedinNewZealand []string `json:"dunedin-new_zealand"`
			GeorgiaUsa        []string `json:"georgia-usa"`
			LosAngelesUsa     []string `json:"los_angeles-usa"`
			NagoyaJapan       []string `json:"nagoya-japan"`
			NorthCarolinaUsa  []string `json:"north_carolina-usa"`
			OsakaJapan        []string `json:"osaka-japan"`
			PenroseNewZealand []string `json:"penrose-new_zealand"`
			SaitamaJapan      []string `json:"saitama-japan"`
		} `json:"datesLocations"`
	} `json:"index"`
}

var Relation Relations
